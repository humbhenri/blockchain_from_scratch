package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/humbhenri/blockchain_from_scratch/blockchain"
	"github.com/humbhenri/blockchain_from_scratch/p2p"
)

var isBootstrapNode bool

func initP2PServer(wg *sync.WaitGroup) {
	defer wg.Done()
	p2p.StartServer()
}

// bootstrapNetwork creates the node network. This node must know who the bootstrap node is, otherwise
// if no bootstrap node is known this node can be one if the flag isBootstrapNode is true
func bootstrapNetwork() (*p2p.Network, error) {
	var network *p2p.Network
	if isBootstrapNode {
		network = p2p.NewNetwork(nil)
		return network, nil
	}
	BLOCKCHAIN_BOOSTRAP_NODE := "BLOCKCHAIN_BOOSTRAP_NODE"
	bootstrapNode := os.Getenv(BLOCKCHAIN_BOOSTRAP_NODE)
	if bootstrapNode == "" {
		return nil, fmt.Errorf("please provide the bootstrap node using the environment variable %s", BLOCKCHAIN_BOOSTRAP_NODE)
	}
	xs := strings.Split(bootstrapNode, " ")
	if len(xs) != 3 {
		return nil, errors.New("bootstrap node must be in format <ID> <IP> <Port>")
	}
	port, err := strconv.Atoi(xs[2])
	if err != nil {
		return nil, fmt.Errorf("port must be a number, but was %s", xs[2])
	}
	node := p2p.Node{ID: xs[0], IP: xs[1], Port: port}
	nodes := []p2p.Node{node}
	return p2p.NewNetwork(nodes), nil
}

func main() {
	flag.BoolVar(&isBootstrapNode, "bootstrap", true, "Indicates that his node is a bootstrap node")
	flag.Parse()

	network, err := bootstrapNetwork()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Initialized network with bootstrap node %v\n", network)
	var wg sync.WaitGroup
	wg.Add(1)
	blockchain.InitBlockChain()
	go initP2PServer(&wg)
	wg.Wait()
}
