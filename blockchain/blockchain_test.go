package blockchain

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestBlockchainCreation(t *testing.T) {
	chain := InitBlockChain(2)
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

func TestMineBlock(t *testing.T) {
	b := genesis(2)
	b.MineBlock(2)
	if !strings.HasPrefix(string(b.Hash), "00") {
		t.Error("Mine failed")
	}
}
