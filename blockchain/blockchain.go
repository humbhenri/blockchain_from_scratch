package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

var nonce int = 0

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
	blocks     []*block
	difficulty int
}

type Blockchain interface {
	AddBlock(string)
	Debug()
	Print(io.Writer)
	IsChainValid() bool
}

var theBlockchain *blockchain

func (b block) String() string {
	return fmt.Sprintf("Block: %s ",
		base64.RawStdEncoding.EncodeToString(b.Hash))
}

func GetBlockchain() Blockchain {
	return theBlockchain
}

func (b *block) deriveHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	nonceStr := strconv.Itoa(nonce)
	info := bytes.Join([][]byte{b.Data, b.PrevHash, timestamp, []byte(nonceStr)}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func createBlock(data string, prevHash []byte, difficulty int) *block {
	block := &block{Hash: []byte{}, Data: []byte(data), PrevHash: prevHash, Timestamp: time.Now().Unix()}
	block.MineBlock(difficulty)
	log.Printf("block mined %v\n", block)
	return block
}

func (chain *blockchain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := createBlock(data, prevBlock.Hash, chain.difficulty)
	chain.blocks = append(chain.blocks, new)
}

func genesis(difficulty int) *block {
	return createBlock("Genesis", []byte{}, difficulty)
}

func InitBlockChain(difficulty int) Blockchain {
	theBlockchain = &blockchain{blocks: []*block{genesis(difficulty)}, difficulty: difficulty}
	return theBlockchain
}

func (chain *blockchain) Debug() {
	for _, block := range chain.blocks {
		fmt.Printf("Block data: %s, hash: %s, previous hash: %s, timestamp: %d\n\n", block.Data,
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
		if slices.Compare(prevBlock.Hash, block.PrevHash) != 0 {
			return false
		}
	}
	return true
}

func (block *block) MineBlock(difficulty int) {
	log.Println("Mining started")
	target := strings.Repeat("0", difficulty)
	for !strings.HasPrefix(string(block.Hash), target) {
		nonce += 1
		block.deriveHash()
	}
	log.Printf("Block mined, nonce = %d\n", nonce)
}
