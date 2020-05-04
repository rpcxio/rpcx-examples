#!/bin/bash

xgen -o cmd/main.go -r "etcd" -pkg github.com/rpcxio/rpcx-examples/xgen

# go run  cmd/main.go