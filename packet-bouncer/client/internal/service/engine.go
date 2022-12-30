package service

import (
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

	for i := 0; i < e.nClients; i++ {
		e.wg.Add(1)
		go e.connWorker(i)

	}
	return nil
}

func (e *Engine) connWorker(num int) {
	conn := client.NewConn(e.connConfig)
	conn.Start()
	defer e.wg.Done()
}
