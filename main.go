package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/humbhenri/blockchain_from_scratch/blockchain"
	"github.com/humbhenri/blockchain_from_scratch/p2p"
)

func initP2PServer(wg *sync.WaitGroup) {
	defer wg.Done()
	p2p.StartServer()
}

func initBootstrap(bootstrapNode string) (*p2p.Network, error) {
	xs := strings.Split(bootstrapNode, " ")
	if len(xs) != 3 {
		return nil, errors.New("bootstrap node must be in format <ID> <IP> <Port>\n")
	}
	port, err := strconv.Atoi(xs[2])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Port must be a number, but was %s\n", xs[2]))
	}
	node := p2p.Node{ID: xs[0], IP: xs[1], Port: port}
	nodes := []p2p.Node{node}
	return p2p.NewNetwork(nodes), nil
}

func main() {
    BLOCKCHAIN_BOOSTRAP_NODE := "BLOCKCHAIN_BOOSTRAP_NODE"
	bootstrapNode := os.Getenv(BLOCKCHAIN_BOOSTRAP_NODE)
	if bootstrapNode == "" {
		fmt.Printf("Please provide the bootstrap node using the environment variable %s\n", BLOCKCHAIN_BOOSTRAP_NODE)
		os.Exit(1)
	}
	network, err := initBootstrap(bootstrapNode)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Initialized network with bootstrap node %v\n", network)
	var wg sync.WaitGroup
	wg.Add(1)
	blockchain.InitBlockChain()
	go initP2PServer(&wg)
	wg.Wait()
}
