#! /bin/sh

version=v1.8.13

go get -v github.com/smallnest/rpcx@$version

cd registry/dynamic_port_allocation
go get -v github.com/smallnest/rpcx@$version
go get -v github.com/rpcxio/rpcx-etcd@HEAD

cd ../../registry/etcd
go get -v github.com/smallnest/rpcx@$version
go get -v github.com/rpcxio/rpcx-etcd@HEAD

cd ../../registry/etcdv3
go get -v github.com/smallnest/rpcx@$version
go get -v github.com/rpcxio/rpcx-etcd@HEAD

cd ../../registry/etcdv3_lazyregister
go get -v github.com/smallnest/rpcx@$version
go get -v github.com/rpcxio/rpcx-etcd@HEAD

cd ../..
