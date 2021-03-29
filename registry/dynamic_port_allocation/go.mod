module github.com/rpcxio/rpcx-examples/registry/dynamic_port_allocation

go 1.15

require (
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/rpcxio/rpcx-etcd v0.0.0-20201229103411-8317fc934fbb
	github.com/rpcxio/rpcx-examples v1.1.6
	github.com/smallnest/rpcx v0.0.0-20210329112732-c584448849f9
	google.golang.org/grpc v1.34.0 // indirect
)

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.3
	google.golang.org/grpc => google.golang.org/grpc v1.29.1
)
