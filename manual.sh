#!/usr/bin/bash

echo "go run ./cmd/web >>./tmp/info.log 2>>./tmp/error.log"
go run ./cmd/web >>./tmp/info.log 2>>./tmp/error.log
