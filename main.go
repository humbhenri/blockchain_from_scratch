package main

import (
	"bytes"
	"crypto/sha256"
	"math/rand"
	"strconv"
	"time"

	"github.com/humbhenri/blockchain_from_scratch/p2p"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	Hash      []byte
	Data      []byte
	PrevHash  []byte
	Timestamp int64
}

func (b *Block) DeriveHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	info := bytes.Join([][]byte{b.Data, b.PrevHash, timestamp}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	n := rand.Intn(3)
	time.Sleep(time.Second * time.Duration(n))
	block := &Block{Hash: []byte{}, Data: []byte(data), PrevHash: prevHash, Timestamp: time.Now().Unix()}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	// chain := InitBlockChain()
	// chain.AddBlock("First after genesis")
	// chain.AddBlock("Second after genesis")
	// chain.AddBlock("Third after genesis")
	// for _, block := range chain.blocks {
	// 	fmt.Printf("Block data: %s, hash: %s, previous hash: %s, timestamp: %d\n", block.Data,
	// 		base64.RawStdEncoding.EncodeToString(block.Hash),
	// 		base64.RawStdEncoding.EncodeToString(block.PrevHash),
	// 		block.Timestamp)
	// }
	p2p.StartServer()
}
