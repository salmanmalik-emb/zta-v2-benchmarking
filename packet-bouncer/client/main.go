package main

import (
	"os"
	"packet-bouncer/client/cmd"
)

func main() {

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
