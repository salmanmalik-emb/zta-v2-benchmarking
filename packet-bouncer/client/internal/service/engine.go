package service

import (
	"encoding/json"
	"fmt"
	"packet-bouncer/client/internal/client"
	"sync"
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
	fmt.Println(string(b))
	return nil
}

func (e *Engine) connWorker(num int, results *[]client.Result) {
	conn := client.NewConn(e.connConfig)
	conn.Start(results)
	defer e.wg.Done()
}
