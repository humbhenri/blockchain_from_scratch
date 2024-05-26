package p2p

import (
	"github.com/google/uuid"
	"strings"
	"errors"
	"fmt"
	"strconv"
)

// Ethereum uses a protocol based on the Kademlia Distributed Hash Table (DHT)
// for node discovery. This protocol helps nodes find each other and is
// essential for maintaining a decentralized network. Here's how it works: Node
// ID and XOR Distance: Each node in the Ethereum network has a unique Node ID,
// which is a 512-bit identifier derived from the node's public key. Nodes use
// the XOR distance metric to find other nodes, where the distance between two
// Node IDs is the XOR of the two IDs. Routing Table: Each node maintains a
// routing table containing information about other nodes. The table is
// structured to facilitate efficient lookups, helping nodes quickly find peers
// whose Node IDs are close to a given target ID.

// Node represents a node in the network
type Node struct {
	ID   uuid.UUID
	IP   string
	Port int
}

// NewNode creates a node from a string
func NewNode(s string) (*Node, error) {
	xs := strings.Split(s, " ")
	if len(xs) != 2 {
		return nil, errors.New("bootstrap node must be in format <IP> <Port>")
	}
	port, err := strconv.Atoi(xs[1])
	if err != nil {
		return nil, fmt.Errorf("port must be a number, but was %s", xs[1])
	}
	return &Node{ID: uuid.New(), IP: xs[0], Port: port}, nil
}


// Network represents the network of nodes
type Network struct {
	Nodes map[uuid.UUID]Node
}

// When a node's routing table is empty, typically when it first joins the
// Ethereum network, it needs to bootstrap the process of finding other nodes.
// This bootstrap process involves several steps:

// Bootstrap List: Each Ethereum client comes with a list of IP addresses and
// ports of these bootstrap nodes.

// Initial Contact: The node sends messages to these bootstrap nodes to announce
// its presence and request information about other active nodes.

// NewNetwork initializes a new network with the bootstrap nodes
func NewNetwork(bootstrap []Node) *Network {
	network := &Network{Nodes: make(map[uuid.UUID]Node)}
	for _, node := range bootstrap {
		network.Nodes[node.ID] = node
	}
	return network
}

// FindNode pings a bootstrap node to stablish a connection
// func FindNode(network *Network) bool {

// }
