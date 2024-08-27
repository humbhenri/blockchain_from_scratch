#!/bin/bash

# This test runs two blockchain nodes, usage: ./test.sh

start_server() {
    PORT=$1
    ./blockchain_from_scratch -port $PORT &
    SERVER_PID=$!
    echo "Started blockchain node with PID $SERVER_PID"
}

send_message_to_server() {
    MSG=$1
    PORT=$2
    echo "$MSG" | nc -q 1 -u 127.0.0.1 $PORT    
}

go build
if [ $? -ne 0 ]; then
    echo "Error: Go build failed!"
    exit 1
fi

start_server 5000
SERVER1_PID=$SERVER_PID

sleep 1

start_server 5001
SERVER2_PID=$SERVER_PID

sleep 1

send_message_to_server "ECHO Hello, Server 1!" 5000
send_message_to_server "ECHO Hello, Server 2!" 5001

sleep 2

kill -9 $SERVER1_PID
kill -9 $SERVER2_PID
