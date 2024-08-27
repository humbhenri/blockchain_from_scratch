package main

import (
	"fmt"

	"github.com/humbhenri/blockchain_from_scratch/server"
)

func processCommands() {
	for {
		// Receive command data from the channel
		cmdData := <-server.DataChannel

		// Process the received command
		switch cmdData.Command {
		case server.Ping:
			fmt.Println("Received PING command with data:", cmdData.Data)
		case server.Echo:
			fmt.Println("Received ECHO command with data:", cmdData.Data)
		case server.Unknown:
			fmt.Println("Received UNKNOWN command with data:", cmdData.Data)
		}
	}
}

func main() {
	go server.StartServer(8080)
	go processCommands()
	select {}
}
