module github.com/rpcxio/rpcx-examples

go 1.16

require (
	github.com/cyberdelia/go-metrics-graphite v0.0.0-20161219230853-39f87cc3b432
	github.com/gogo/protobuf v1.3.1
	github.com/json-iterator/go v1.1.10
	github.com/kr/pretty v0.2.1
	github.com/nacos-group/nacos-sdk-go v1.0.7
	github.com/opentracing/opentracing-go v1.2.0
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/rpcxio/rpcx-nacos v0.0.0-20210525063414-e99e0c81ae30
	github.com/smallnest/rpcx v1.6.12-0.20211102092848-d3baafd4aa3c
	github.com/xtaci/kcp-go v5.4.20+incompatible
	go.opencensus.io v0.22.3 // indirect
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	google.golang.org/grpc/examples v0.0.0-20211005235525-4bd99953513f // indirect
)

// replace github.com/smallnest/rpcx => ../../smallnest/rpcx
