module github.com/rpcxio/rpcx-examples/registry/dynamic_port_allocation

go 1.15

require (
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/rpcxio/rpcx-etcd v0.0.0-20201223114009-7e3fe9c7ae7e
	github.com/rpcxio/rpcx-examples v1.1.6
	github.com/smallnest/rpcx v0.0.0-20201223122003-ca98a3ecb90c
)

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.3
	google.golang.org/grpc => google.golang.org/grpc v1.29.1
)