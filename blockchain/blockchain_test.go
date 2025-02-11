package blockchain

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestBlockCreation(t *testing.T)  {
	block := createBlock("Genesis", nil)
	if block.Hash == nil {
		t.Error("hash not be empty")
	}

	block2 := createBlock("Hello world", block.Hash)
	if block2.Hash == nil {
		t.Error("hash not be empty")
	}

}

func TestBlockchainCreation(t *testing.T) {
	chain := InitBlockChain()
	chain.AddBlock("Data 1")
	chain.AddBlock("Data 2")
	chain.AddBlock("Data 3")
	var buf bytes.Buffer
	chain.Print(io.MultiWriter(&buf, os.Stderr))
	if buf.Len() == 0 {
		t.Error("Output should not be empty")
	}

	if !chain.IsChainValid() {
		t.Error("Blockchain is not valid.")
	}
}
