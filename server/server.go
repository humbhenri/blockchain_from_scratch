package server

import (
	"log"
	"net"
	"strconv"
	"strings"
)

// Command represents different types of commands
type Command int

const (
	Unknown Command = iota
	Ping
	Echo
	AddData
	Print
)

// CommandNames maps Command values to their names
var CommandNames = map[Command]string{
	Unknown: "UNKNOWN",
	Ping:    "PING",
	Echo:    "ECHO",
	AddData: "ADD_DATA",
	Print:   "PRINT",
}

// CommandData holds the command and its associated data
type CommandData struct {
	Command Command
	Data    string
}

// DataChannel is a channel for sending commands and their data
var DataChannel = make(chan CommandData)

func StartServer(port int) {
	portStr := strconv.Itoa(port)

	// Resolve the UDP address to listen on
	addr, err := net.ResolveUDPAddr("udp", ":"+portStr)
	if err != nil {
		log.Fatalf("Error resolving address: %s\n", err)
	}

	// Create a UDP connection for listening
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Error listening on UDP: %s\n", err)
	}
	defer conn.Close()

	log.Printf("Node listening on UDP port %d\n", port)

	// Buffer to store incoming data
	buf := make([]byte, 1024)

	for {
		// Read from the UDP connection
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Error reading from UDP:", err)
			continue
		}

		message := string(buf[:n])
		cmd, data := parseCommand(message)

		// Send the command and its data to the channel
		DataChannel <- CommandData{Command: cmd, Data: data}
	}
}

// parseCommand parses the message and returns the command and its data
func parseCommand(message string) (Command, string) {
	parts := strings.SplitN(message, " ", 2)
	if len(parts) < 1 {
		return Unknown, message
	}

	switch strings.ToUpper(parts[0]) {
	case "PING":
		return Ping, getData(parts)
	case "ECHO":
		return Echo, getData(parts)
	case "ADD_DATA":
		return AddData, getData(parts)
	case "PRINT":
		return Print, getData(parts)
	default:
		return Unknown, message
	}
}

// getData returns the data part of the command, or an empty string if not present
func getData(parts []string) string {
	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}
