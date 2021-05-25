#! /bin/sh

version=HEAD

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
