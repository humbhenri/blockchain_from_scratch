// handles communication between peers

package p2p

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/humbhenri/blockchain_from_scratch/blockchain"
)

const (
	INITIAL_PORT = 33033
	LIMIT_PORT   = 33100
)

// StartServer starts a server to handle incoming p2p requests, it first tries to bind to a fixed port,
// if that results in an error it tries to bind to next until a port limit is reached.
// This function blocks.
func StartServer() {
	port := INITIAL_PORT
	for port < LIMIT_PORT {
		log.Printf("Listening in port %d\n", port)
		err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
		log.Printf("%v\n", err)
		if strings.Contains(err.Error(), "address already in use") {
			port++
		} else {
			log.Fatal(err)
		}
	}
	log.Fatal("Failt to start the server")
}
