package client

import (
	"encoding/binary"
	"net"
	"os"
	"sync"
	"time"
)

type ConnConfig struct {
	Endpoint          string
	Duration          int
	PacketSize        int
	Pps               int
	Protocol          string
	StopDelayDuration int
}

type clientConn struct {
	config ConnConfig
	c      net.Conn
	mutex  sync.Mutex
}

type rtt struct {
	Rtts0 float64 `json:"0_49"`
	Rtts1 float64 `json:"50_179"`
	Rtts2 float64 `json:"180_399"`
	Rtts3 float64 `json:"400_999"`
	Rtts4 float64 `json:"1000_1999"`
	Rtts5 float64 `json:"2000+"`
}

type Result struct {
	ReceivedPackets int     `json:"received_packets"`
	MinMaxWindow    int     `json:"min_max_window"`
	SentPackets     int     `json:"sent_packets"`
	Loss            float64 `json:"loss"`
	Badl            float64 `json:"bad_loss"`
	Rtt             rtt
	Score           float64 `json:"score"`
	Throughput      float64 `json:"throughput_Mbs"`
}

func NewConn(config ConnConfig) *clientConn {

	//c, err := net.Dial("udp", config.Endpoint)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil
	//}
	return &clientConn{
		config: config,
		c:      nil,
	}
}

func (conn *clientConn) Start(results *[]Result) error {

	var c net.Conn
	if conn.config.Protocol == "udp" {
		s, err := net.ResolveUDPAddr(conn.config.Protocol, conn.config.Endpoint)

		if err != nil {
			println("ResolveUDPAddr failed:", err.Error())
			os.Exit(1)
		}

		c, err = net.DialUDP(conn.config.Protocol, nil, s)
		if err != nil {
			println("Listen failed:", err.Error())
			os.Exit(1)
		}
	} else if conn.config.Protocol == "tcp" {
		s, err := net.ResolveTCPAddr(conn.config.Protocol, conn.config.Endpoint)

		if err != nil {
			println("ResolveUDPAddr failed:", err.Error())
			os.Exit(1)
		}

		c, err = net.DialTCP(conn.config.Protocol, nil, s)
		if err != nil {
			println("Listen failed:", err.Error())
			os.Exit(1)
		}
	}

	//close the connection
	defer c.Close()

	timeBase := time.Now()

	wg := sync.WaitGroup{}

	wg.Add(2)
	go conn.sendingThread(c, timeBase, &wg)
	go conn.receivingThread(c, timeBase, &wg, results)

	wg.Wait()
	return nil
}

func (conn *clientConn) sendingThread(c net.Conn, timeBase time.Time, wg *sync.WaitGroup) {
	time.Sleep(500 * time.Millisecond) // to allow receiver to warm up
	start := time.Now()
	step := time.Second / time.Duration(conn.config.Pps)
	n := conn.config.Pps * int(conn.config.Duration)

	buf := make([]byte, conn.config.PacketSize)

	var totalctr uint32
	totalctr = 0

	for i := 0; i < n; i++ {
		deadline := start.Add(step * time.Duration(i))
		now := time.Now()

		var delta time.Duration
		if now.Before(deadline) {
			delta = deadline.Sub(timeBase)
			time.Sleep(deadline.Sub(now))
		} else {
			delta = now.Sub(timeBase)
		}

		//binary.BigEndian.PutUint64(buf[0:8], uint64(delta.Seconds()))
		binary.BigEndian.PutUint64(buf[0:8], uint64(delta.Nanoseconds()))

		binary.BigEndian.PutUint32(buf[8:12], totalctr)

		//fmt.Println(fmt.Sprintf("write s=%d, ns=%d, no=%d", uint64(delta.Seconds()), uint32(delta.Nanoseconds()), totalctr))
		_, err := c.Write(buf)
		if err == nil {
			totalctr += 1
		}
	}
	//fmt.Println("packets sent: ", totalctr)
	wg.Done()
}

type Packet struct {
	no   uint32
	rtt4 time.Duration
}

func (p Packet) Less(other Packet) bool {
	return p.no > other.no
}

func (conn *clientConn) receivingThread(c net.Conn, timeBase time.Time, wg *sync.WaitGroup, results *[]Result) {
	jitterBuffer := make([]Packet, 0, 1024)

	deadline := time.Now().Add(time.Duration(conn.config.Duration+conn.config.StopDelayDuration) * time.Second)

	minN := ^uint32(0)
	maxN := uint32(0)

	ctr := 0
	prevn := ^uint32(0)
	badloss := 0
	dup := 0
	rtt4stats := [6]int{}

	receivedSomething := false

	analysePacket := func(p *Packet) {
		if minN > p.no {
			minN = p.no
		}
		if maxN < p.no {
			maxN = p.no
		}

		if p.no == prevn {
			// duplicate packet, ignore
			dup += 1
			return
		}

		if p.no > prevn && p.no-prevn > 50 {
			badloss += int(p.no - prevn)
		}

		if p.rtt4 < 50*time.Millisecond {
			rtt4stats[0] += 1
		} else if p.rtt4 < 180*time.Millisecond {
			rtt4stats[1] += 1
		} else if p.rtt4 < 400*time.Millisecond {
			rtt4stats[2] += 1
		} else if p.rtt4 < 1000*time.Millisecond {
			rtt4stats[3] += 1
		} else if p.rtt4 < 2000*time.Millisecond {
			rtt4stats[4] += 1
		} else {
			rtt4stats[5] += 1
		}

		ctr += 1
		prevn = p.no
	}

	buf := make([]byte, conn.config.PacketSize)
	receiveStartTime := time.Now()
	var no uint32
	for {
		c.SetReadDeadline(deadline)
		_, err := c.Read(buf)
		now := time.Now()
		if now.After(deadline) {
			break
		}
		if err == nil {
			//println("Received a packet from ", addr)

			//s := binary.BigEndian.Uint64(buf[:8])
			ns := binary.BigEndian.Uint64(buf[0:8])
			no = binary.BigEndian.Uint32(buf[8:12])

			if /*s == 0 &&*/ ns == 0 && no == 0 {
				continue
			}
			sinceBase := time.Duration(ns) * time.Nanosecond
			sendTime := timeBase.Add(sinceBase)
			var rtt4 time.Duration
			if now.After(sendTime) {
				rtt4 = now.Sub(sendTime)
			} else {
				rtt4 = time.Duration(999) * time.Second
			}

			//fmt.Println(fmt.Sprintf("read ns=%d, no=%d, rtt=%v, sendTime=%v, now=%v", ns, no, rtt4, sendTime, now))
			jitterBuffer = append(jitterBuffer, Packet{
				no:   no,
				rtt4: rtt4,
			})
			if len(jitterBuffer) > 1022 {
				analysePacket(&jitterBuffer[len(jitterBuffer)-1])
				jitterBuffer = jitterBuffer[:len(jitterBuffer)-1]
			}

			if !receivedSomething {
				receivedSomething = true
				// log message
			}
		} else {
			break
			// ignore error
		}
	}
	for _, p := range jitterBuffer {
		analysePacket(&p)
	}
	receiveCumulativeTimeDelta := time.Now().Sub(receiveStartTime)
	//receiveCumulativeTimeDelta -= time.Duration(conn.config.StopDelayDuration)
	if ctr == 0 {
		return
	}
	nn := float64(maxN - minN + 1)
	loss := 100.0 - float64(ctr)*100.0/nn
	badl := float64(badloss) * 100.0 / nn
	rtts0 := float64(rtt4stats[0]) * 100.0 / nn
	rtts1 := float64(rtt4stats[1]) * 100.0 / nn
	rtts2 := float64(rtt4stats[2]) * 100.0 / nn
	rtts3 := float64(rtt4stats[3]) * 100.0 / nn
	rtts4 := float64(rtt4stats[4]) * 100.0 / nn
	rtts5 := float64(rtt4stats[5]) * 100.0 / nn
	score := 10.0 - loss/3.0 - badl/1.0 - rtts2/200.0 - rtts3/80.0 - rtts4/40.0 - rtts5/20.0
	//fmt.Println(receiveCumulativeTimeDelta.Seconds(), " ctr", ctr)
	//fmt.Println("ctr: ", ctr, "packetsize", conn.config.PacketSize, "timedelta", receiveCumulativeTimeDelta.Seconds())
	throughput := ((float64(ctr) * float64(conn.config.PacketSize+40)) / receiveCumulativeTimeDelta.Seconds()) / float64(1024*1024)
	//fmt.Println(throughput * 60 * 2)
	rtt := rtt{
		Rtts0: rtts0,
		Rtts1: rtts1,
		Rtts2: rtts2,
		Rtts3: rtts3,
		Rtts4: rtts4,
		Rtts5: rtts5,
	}

	result := Result{
		Loss:            loss,
		Badl:            badl,
		Rtt:             rtt,
		Score:           score,
		Throughput:      throughput * 8 * 2,
		ReceivedPackets: ctr,
		SentPackets:     conn.config.Pps * conn.config.Duration,
	}
	conn.mutex.Lock()
	*results = append(*results, result)
	conn.mutex.Unlock()

	//fmt.Println(fmt.Sprintf("%+v", result))
	//fmt.Println("Throughput: ", throughput*8*2, " Mb/s")

	wg.Done()
}
