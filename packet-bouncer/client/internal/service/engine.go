package service

import (
	"encoding/json"
	"fmt"
	"packet-bouncer/client/internal/client"
	"sync"
	"time"
)

type Engine struct {
	connConfig client.ConnConfig
	nClients   int
	wg         *sync.WaitGroup
}

// NewEngine creates a new Connection Engine
func NewEngine(
	connConfig client.ConnConfig,
	nClients int,
	wg *sync.WaitGroup,
) *Engine {
	return &Engine{
		connConfig: connConfig,
		nClients:   nClients,
		wg:         wg,
	}
}

func (e *Engine) Stop() error {
	return nil
}

func (e *Engine) Start() error {
	var results []client.Result
	for i := 0; i < e.nClients; i++ {
		e.wg.Add(1)
		go e.connWorker(i, &results)
	}
	fmt.Println("Test started, start time: ", time.Now().String())
	approxThroughput := (float64(e.connConfig.Pps*e.nClients*(e.connConfig.PacketSize+40)*2*8) / float64(1024*1024) / 8) * 60
	fmt.Println("This script will produce approximately: ", approxThroughput, " Mb/Minute of throughput")
	e.wg.Wait()
	totalThreads := len(results)
	finalResult := client.Result{}
	for _, result := range results {
		finalResult.ReceivedPackets += result.ReceivedPackets
		finalResult.Loss += result.Loss
		finalResult.Badl += result.Badl
		finalResult.Rtt.Rtts0 += result.Rtt.Rtts0
		finalResult.Rtt.Rtts1 += result.Rtt.Rtts1
		finalResult.Rtt.Rtts2 += result.Rtt.Rtts2
		finalResult.Rtt.Rtts3 += result.Rtt.Rtts3
		finalResult.Rtt.Rtts4 += result.Rtt.Rtts4
		finalResult.Rtt.Rtts5 += result.Rtt.Rtts5
		finalResult.Score += result.Score
		finalResult.SentPackets += result.SentPackets
		finalResult.Throughput += result.Throughput
		finalResult.Status += result.Status + "\n"
	}

	//finalResult.ReceivedPackets = finalResult.ReceivedPackets / totalThreads
	finalResult.Loss = finalResult.Loss / float64(totalThreads)
	finalResult.Badl = finalResult.Badl / float64(totalThreads)
	finalResult.Rtt.Rtts0 = finalResult.Rtt.Rtts0 / float64(totalThreads)
	finalResult.Rtt.Rtts1 = finalResult.Rtt.Rtts1 / float64(totalThreads)
	finalResult.Rtt.Rtts2 = finalResult.Rtt.Rtts2 / float64(totalThreads)
	finalResult.Rtt.Rtts3 = finalResult.Rtt.Rtts3 / float64(totalThreads)
	finalResult.Rtt.Rtts4 = finalResult.Rtt.Rtts4 / float64(totalThreads)
	finalResult.Rtt.Rtts5 = finalResult.Rtt.Rtts5 / float64(totalThreads)
	finalResult.Score = finalResult.Score / float64(totalThreads)
	//finalResult.SentPackets = finalResult.SentPackets / totalThreads
	//finalResult.Throughput = finalResult.Throughput / float64(totalThreads)
	b, err := json.Marshal(finalResult)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Test End Time:", time.Now().String())
	fmt.Println("Approx Throughput in MB/minute  (aws format): ", (finalResult.Throughput/8.0)*60)
	fmt.Println(string(b))
	return nil
}

func (e *Engine) connWorker(num int, results *[]client.Result) {
	conn := client.NewConn(e.connConfig)
	conn.Start(results)
	defer e.wg.Done()
}
