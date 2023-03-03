package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"packet-bouncer/client/internal/client"
	"packet-bouncer/client/internal/service"
	"sync"
)

var (
	endpoint          string
	nClients          int
	pps               int
	packetSize        int
	duration          int
	protocol          string
	engine            *service.Engine
	stopDelayDuration int
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start service",
	RunE: func(cmd *cobra.Command, args []string) error {

		conf := client.ConnConfig{
			Pps:               pps,
			Endpoint:          endpoint,
			Duration:          duration,
			PacketSize:        packetSize,
			Protocol:          protocol,
			StopDelayDuration: stopDelayDuration,
		}

		wg := sync.WaitGroup{}

		engine = service.NewEngine(conf, nClients, &wg)

		err := engine.Start()

		if err != nil {
			cmd.Printf("Failed to start Client Service")
			return err
		}

		wg.Wait()

		err = engine.Stop()
		if err != nil {
			fmt.Println(fmt.Printf("failed stopping engine %v", err))
			return err
		}

		return nil
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop service",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := engine.Stop()
		if err != nil {
			cmd.Printf("Failed to stop Service")
			return err
		}
		cmd.Printf("Service has been stopped")
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&endpoint, "endpoint", "localhost:8050", "sets endpoint")
	rootCmd.PersistentFlags().IntVar(&nClients, "clients", 10, "sets no. of clients")
	rootCmd.PersistentFlags().IntVar(&duration, "duration", 6000, "sets duration")
	rootCmd.PersistentFlags().IntVar(&stopDelayDuration, "delay-after-stopping-sender", 300, "seconds to wait and receive after stopping sender")
	rootCmd.PersistentFlags().IntVar(&pps, "pps", 90, "sets packets per seconds for each client")
	rootCmd.PersistentFlags().IntVar(&packetSize, "packet", 960, "sets packet size for each packet")
	rootCmd.PersistentFlags().StringVar(&protocol, "protocol", "tcp", "sets protocol")
}

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "manages ZT Client service",
}
