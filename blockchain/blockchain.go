package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
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
	PrevHash  string `json:"prev_hash"`
	Data      string `json:"data"`
	Timestamp int64  `json:"timestamp"`
}

// TimestampFmt returns the timestamp formatted for humans
func (b *JsonBlock) TimestampFmt() string  {
    t := time.Unix(b.Timestamp, 0)
    return t.Format("02/01/2006, 15:04:05")
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
    SetDifficulty(int)
    Blocks() []*JsonBlock
}

var theBlockchain *blockchain

func (b block) String() string {
	return fmt.Sprintf("Block: %s ", hex.EncodeToString(b.Hash))
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
			hex.EncodeToString(block.Hash),
			hex.EncodeToString(block.PrevHash),
			block.Timestamp)
	}
}

func (chain *blockchain) Print(writer io.Writer) {
	enc := json.NewEncoder(writer)
	enc.SetIndent("", "  ")
	var blocks []JsonBlock
	for _, block := range chain.blocks {
		blocks = append(blocks, JsonBlock{
			Hash:      hex.EncodeToString(block.Hash),
			PrevHash:  hex.EncodeToString(block.PrevHash),
			Data:      string(block.Data),
			Timestamp: block.Timestamp,
		})
	}
	enc.Encode(blocks)
}

func (chain *blockchain) Blocks() []*JsonBlock {
    var blocks []*JsonBlock
	for _, block := range chain.blocks {
		blocks = append(blocks, &JsonBlock{
			Hash:      hex.EncodeToString(block.Hash),
			PrevHash:  hex.EncodeToString(block.PrevHash),
			Data:      string(block.Data),
			Timestamp: block.Timestamp,
		})
	}
    return blocks
}

func Load(reader io.Reader) error {
	enc := json.NewDecoder(reader)
	var json_blocks []JsonBlock
	err := enc.Decode(&json_blocks)
	if err != nil {
		return err
	}
	var blocks []*block
	for _, b := range json_blocks {
		hash, err := hex.DecodeString(b.Hash)
		if err != nil {
			return err
		}
		prevHash, err := hex.DecodeString(b.PrevHash)
        if err != nil {
            return err
        }
		blocks = append(blocks, &block{
			Hash:      hash,
			PrevHash:  prevHash,
			Timestamp: b.Timestamp,
			Data:      []byte(b.Data),
		})
	}
    theBlockchain = &blockchain{blocks: blocks}
	return nil
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

func (chain *blockchain) SetDifficulty(difficulty int) {
    chain.difficulty = difficulty
}
