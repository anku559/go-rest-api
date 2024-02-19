#!/bin/bash

# /home/ubuntu/go/bin/CompileDaemon -directory=$PWD -build="go build" -command="go run ." -color=true -log-prefix=true

nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run $PWD/cmd/gateway/main.go