package main

import (
	"os"
	"github.com/seafooler/consensusPBFT/pbft/network"
)

func main() {
	nodeID := os.Args[1]
	server := network.NewServer(nodeID)

	server.Start()
}
