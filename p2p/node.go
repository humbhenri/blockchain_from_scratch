package p2p

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
	ID string
	IP string
	Port int
}

// Network represents the network of nodes
type Network struct {
	Nodes map[string]Node
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
	network := &Network{Nodes: make(map[string]Node)}
	for _, node := range bootstrap {
		network.Nodes[node.ID] = node
	}
	return network
}

