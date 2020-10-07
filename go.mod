module github.com/rpcxio/rpcx-examples

go 1.15

require (
	github.com/cyberdelia/go-metrics-graphite v0.0.0-20161219230853-39f87cc3b432
	github.com/gogo/protobuf v1.3.1
	github.com/json-iterator/go v1.1.10
	github.com/kr/pretty v0.2.1
	github.com/nacos-group/nacos-sdk-go v1.0.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/rpcxio/rpcx-gateway v0.0.0-20200521025828-a39934d3752d
	github.com/smallnest/rpcx v0.0.0-20200924044220-f2cdd4dea15a
	github.com/xtaci/kcp-go v5.4.20+incompatible
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.0
