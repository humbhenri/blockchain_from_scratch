package main

import (
	"github.com/humbhenri/blockchain_from_scratch/p2p"
	"github.com/humbhenri/blockchain_from_scratch/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.Debug()
	p2p.StartServer()
}
