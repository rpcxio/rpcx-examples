#! /bin/sh

version=475481d3b37b454b60014a09b5202f801c1c4291

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
