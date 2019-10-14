#!/bin/bash

xgen -o cmd/main.go -r "etcd" -pkg github.com/rpcx-ecosystem/rpcx-examples3/xgen

# go run  cmd/main.go