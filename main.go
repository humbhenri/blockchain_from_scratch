package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/humbhenri/blockchain_from_scratch/blockchain"
	"github.com/humbhenri/blockchain_from_scratch/server"
)

func processCommands() {
	chain := blockchain.GetBlockchain()
	ticker := time.NewTicker(15 * time.Second)
	go func() {
		for {
			select {
			case cmdData := <-server.DataChannel:
				// Process the received command
				switch cmdData.Command {
				case server.Ping:
					log.Println("Received PING command with data:", cmdData.Data)
				case server.Echo:
					log.Println("Received ECHO command with data:", cmdData.Data)
				case server.AddData:
					chain.AddBlock(cmdData.Data)
				case server.Print:
					chain.Debug()
				case server.Unknown:
					log.Println("Received UNKNOWN command with data:", cmdData.Data)
				}
			case <-ticker.C:
				log.Println("Saving blockchain in file system ...")
				f, err := os.OpenFile("/tmp/block.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
				if err != nil {
					panic(err)
				}
				defer f.Close()
				chain.Print(f)
			}
		}
	}()
}

func main() {
	port := flag.Int("port", 8080, "UDP port to listen on")
	difficulty := flag.Int("difficulty", 2, "proof of work difficulty")
	flag.Parse()

	blockchain.InitBlockChain(*difficulty)
	go server.StartServer(*port)
	processCommands()
	log.Println("Blockchain started")

	select {}
}
