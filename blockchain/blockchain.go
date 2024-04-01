package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"encoding/base64"
)

type block struct {
	Hash      []byte
	Data      []byte
	PrevHash  []byte
	Timestamp int64
}

type blockchain struct {
	blocks []*block
}

type Blockchain interface {
	AddBlock(string)
	Debug()
}

func (b *block) deriveHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	info := bytes.Join([][]byte{b.Data, b.PrevHash, timestamp}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func createBlock(data string, prevHash []byte) *block {
	n := rand.Intn(3)
	time.Sleep(time.Second * time.Duration(n))
	block := &block{Hash: []byte{}, Data: []byte(data), PrevHash: prevHash, Timestamp: time.Now().Unix()}
	block.deriveHash()
	return block
}

func (chain *blockchain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := createBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func genesis() *block {
	return createBlock("Genesis", []byte{})
}

func InitBlockChain() Blockchain {
	return &blockchain{[]*block{genesis()}}
}

func (chain *blockchain) Debug() {
	for _, block := range chain.blocks {
		fmt.Printf("Block data: %s, hash: %s, previous hash: %s, timestamp: %d\n", block.Data,
			base64.RawStdEncoding.EncodeToString(block.Hash),
			base64.RawStdEncoding.EncodeToString(block.PrevHash),
			block.Timestamp)
	}
}
