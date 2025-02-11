package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/exp/slices"
)

type block struct {
	Hash      []byte
	Data      []byte
	PrevHash  []byte
	Timestamp int64
}

type JsonBlock struct {
	Hash      string `json:"hash"`
	Data      string `json:"data"`
	Timestamp int64  `json:"timestamp"`
}

type blockchain struct {
	blocks []*block
}

type Blockchain interface {
	AddBlock(string)
	Debug()
	Print(io.Writer)
	IsChainValid() bool
}

var theBlockchain *blockchain

func GetBlockchain() Blockchain {
	return theBlockchain
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
	theBlockchain = &blockchain{[]*block{genesis()}}
	return theBlockchain
}

func (chain *blockchain) Debug() {
	for _, block := range chain.blocks {
		fmt.Printf("Block data: %s, hash: %s, previous hash: %s, timestamp: %d\n", block.Data,
			base64.RawStdEncoding.EncodeToString(block.Hash),
			base64.RawStdEncoding.EncodeToString(block.PrevHash),
			block.Timestamp)
	}
}

func (chain *blockchain) Print(writer io.Writer) {
	enc := json.NewEncoder(writer)
	enc.SetIndent("", "  ")
	var blocks []JsonBlock
	for _, block := range chain.blocks {
		blocks = append(blocks, JsonBlock{
			Hash:      base64.RawStdEncoding.EncodeToString(block.Hash),
			Data:      string(block.Data),
			Timestamp: block.Timestamp,
		})
	}
	enc.Encode(blocks)
}

// IsChainValid test for blockchain integrity
func (chain *blockchain) IsChainValid() bool {
	for i := 1; i < len(chain.blocks); i++ {
		prevBlock := chain.blocks[i-1]
		block := chain.blocks[i]
		if slices.Compare(prevBlock.Hash,  block.PrevHash) != 0 {
			return false
		}
	}
	return true
}
