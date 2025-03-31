#!/bin/env sh

CMD=$1
DATA=$2
PORT=8080

echo -n "$CMD $DATA" | nc -4u -w0 localhost $PORT
