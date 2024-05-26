package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/humbhenri/blockchain_from_scratch/blockchain"
	"github.com/humbhenri/blockchain_from_scratch/p2p"
)

var bootstrapNode string

func initP2PServer(wg *sync.WaitGroup) {
	defer wg.Done()
	p2p.StartServer()
}

// bootstrapNetwork creates the node network. This node must know who the bootstrap node is, otherwise
// if no bootstrap node is known this node can be one if the flag isBootstrapNode is true
func bootstrapNetwork() (*p2p.Network, error) {
	var network *p2p.Network
	if bootstrapNode == "" {
		network = p2p.NewNetwork(nil)
		return network, nil
	}
	node, err := p2p.NewNode(bootstrapNode)
	if err != nil {
		return nil, err
	}
	nodes := []p2p.Node{*node}
	return p2p.NewNetwork(nodes), nil
}

func main() {
	flag.StringVar(&bootstrapNode, "bootstrap", "", "The bootstrap node to connect to, if empty this node is the bootstrap node")
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
