#!/bin/sh

if [ -z $1 ]; then
    echo "cmd is required"
    exit 1
fi

CMD=$1
DATA=$2
PORT="${3:-8080}"

echo "$CMD $DATA" | nc localhost $PORT
