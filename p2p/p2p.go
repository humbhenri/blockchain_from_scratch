// handles communication between peers

package p2p

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/humbhenri/blockchain_from_scratch/blockchain"
)

const (
	INITIAL_PORT = 33033
	LIMIT_PORT   = 33100
)

// handleGetBlocks print the blocks information from the genesis til the end in json format
func handleGetBlocks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	blockchain.GetBlockchain().Print(w)
}

// handleGetBlocks print the blocks information from the genesis til the end in json format
func handleCreateBlock(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading request body: %v", err)
		return
	}

	fmt.Printf("Received POST request with body: %s\n", body)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Received POST request with body: %s", body)
	blockchain.GetBlockchain().AddBlock(string(body))
}

// StartServer starts a server to handle incoming p2p requests, it first tries to bind to a fixed port,
// if that results in an error it tries to bind to next until a port limit is reached.
// This function blocks.
func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/blocks", handleGetBlocks).Methods("GET")
	router.HandleFunc("/blocks", handleCreateBlock).Methods("POST")

	server := &http.Server{
		Handler: router,
	}
	port := INITIAL_PORT
	for port < LIMIT_PORT {
		log.Printf("Listening in port %d\n", port)
		server.Addr = ":" + strconv.Itoa(port)
		err := server.ListenAndServe()
		log.Printf("%v\n", err)
		if strings.Contains(err.Error(), "address already in use") {
			port++
		} else {
			log.Fatal(err)
		}
	}
	log.Fatal("Failt to start the server")
}
