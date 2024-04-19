package p2p

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/humbhenri/blockchain_from_scratch/blockchain"
)

func TestCreateBlock(t *testing.T) {
	blockchain.InitBlockChain()

	req, err := http.NewRequest("POST", "/blocks", strings.NewReader("teste"))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handleCreateBlock)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Received POST request with body: teste"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
