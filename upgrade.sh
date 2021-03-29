#! /bin/sh

version=c584448849f9173a81d8d0a8a0e305bffd920629

go get -v github.com/smallnest/rpcx@$version

cd registry/dynamic_port_allocation
go get -v github.com/smallnest/rpcx@$version

cd ../../registry/etcd
go get -v github.com/smallnest/rpcx@$version

cd ../../registry/etcdv3
go get -v github.com/smallnest/rpcx@$version

cd ../../registry/etcdv3_lazyregister
go get -v github.com/smallnest/rpcx@$version

cd ../..
