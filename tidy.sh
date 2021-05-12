#! /bin/sh

version=master

go mod tidy
go mod download

cd registry/dynamic_port_allocation
go mod tidy
go mod download

cd ../../registry/etcd
go mod tidy
go mod download

cd ../../registry/etcdv3
go mod tidy
go mod download

cd ../../registry/etcdv3_lazyregister
go mod tidy
go mod download

cd ../..
