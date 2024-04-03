package main

import (
	"github.com/humbhenri/blockchain_from_scratch/blockchain"
	"github.com/humbhenri/blockchain_from_scratch/p2p"
)

func main() {
	blockchain.InitBlockChain()
	p2p.StartServer()
}
