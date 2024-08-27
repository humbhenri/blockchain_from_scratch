package main

import (
	"flag"
	"fmt"

	"github.com/humbhenri/blockchain_from_scratch/blockchain"
	"github.com/humbhenri/blockchain_from_scratch/server"
)

func processCommands() {
	chain := blockchain.GetBlockchain()
	for {
		// Receive command data from the channel
		cmdData := <-server.DataChannel

		// Process the received command
		switch cmdData.Command {
		case server.Ping:
			fmt.Println("Received PING command with data:", cmdData.Data)
		case server.Echo:
			fmt.Println("Received ECHO command with data:", cmdData.Data)
		case server.AddData:
			chain.AddBlock(cmdData.Data)
			fmt.Println("Add block")
			chain.Debug()
		case server.Unknown:
			fmt.Println("Received UNKNOWN command with data:", cmdData.Data)
		}
	}
}

func main() {
	port := flag.Int("port", 8080, "UDP port to listen on")
	blockchain.InitBlockChain()
	flag.Parse()
	go server.StartServer(*port)
	go processCommands()
	fmt.Println("Blockchain started")
	select {}
}
