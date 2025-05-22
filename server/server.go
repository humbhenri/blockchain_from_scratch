package server

import (
	"bufio"
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

	ln, err := net.Listen("tcp", ":"+portStr)
	if err != nil {
		log.Fatalf("Error resolving address: %s\n", err)
	}
	defer ln.Close()
	log.Printf("Listening on %s", portStr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Error listening on TCP: %s\n", err)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading: %v", err)
			conn.Close()
			return
		}
		log.Printf("Received %s", message)
		cmd, data := parseCommand(message)
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
