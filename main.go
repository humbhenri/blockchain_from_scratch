package main

import (
	"flag"
	"html/template"
	"log"
	//	"strconv"
	"time"

	"github.com/humbhenri/blockchain_from_scratch/blockchain"
	"github.com/humbhenri/blockchain_from_scratch/fs"
	"github.com/humbhenri/blockchain_from_scratch/server"

	"net/http"
)

func processCommands(port int) {
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
				log.Println("Saving data to fs ...")
				w := fs.OutputStream(port)
				chain.Print(w)
				w.Close()
			}
		}
	}()
}

var tpl *template.Template

func main() {
	port := flag.Int("port", 8080, "TCP port to listen on")
	difficulty := flag.Int("difficulty", 2, "proof of work difficulty")
	flag.Parse()

	r := fs.ReadStream(*port)
	err := blockchain.Load(r)
	if err != nil {
		log.Println("blockchain read error", err)
        blockchain.InitBlockChain(*difficulty)
	} else {
		blockchain.GetBlockchain().SetDifficulty(*difficulty)
	}
	go server.StartServer(*port)
	processCommands(*port)
	log.Println("Blockchain started")
	blockchain.GetBlockchain().Debug()

	select {}
    // tpl, _ = template.ParseGlob("templates/*.html")
    // http.HandleFunc("/", BlockExplorerHandleFunc)
    // portStr := strconv.Itoa(*port)
    // log.Printf("Block explorer listening on http://localhost:%s\n", portStr)
    // log.Fatal(http.ListenAndServe(":" + portStr, nil))
}

func BlockExplorerHandleFunc(w http.ResponseWriter, r *http.Request) {
    blocks := blockchain.GetBlockchain().Blocks()
    tpl.ExecuteTemplate(w, "index.html", blocks)
}
