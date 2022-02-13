module github.com/rpcxio/rpcx-examples

go 1.16

require (
	github.com/cyberdelia/go-metrics-graphite v0.0.0-20161219230853-39f87cc3b432
	github.com/gogo/protobuf v1.3.2
	github.com/json-iterator/go v1.1.12
	github.com/nacos-group/nacos-sdk-go v1.0.7
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475
	github.com/rpcxio/rpcx-nacos v0.0.0-20210525063414-e99e0c81ae30
	github.com/smallnest/rpcx v1.7.3
	github.com/xtaci/kcp-go v5.4.20+incompatible
	go.opentelemetry.io/otel v1.3.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.3.0
	go.opentelemetry.io/otel/sdk v1.3.0
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3
)

require github.com/kr/pretty v0.2.1 // indirect

// replace github.com/smallnest/rpcx => ../../smallnest/rpcx
