package main

import (
	"flag"
	"log"

	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	example "github.com/rpcxio/rpcx-examples"
	nserverplugin "github.com/rpcxio/rpcx-nacos/serverplugin"
	"github.com/smallnest/rpcx/server"
)

var addr = flag.String("addr", "127.0.0.1:8972", "server address")

func main() {
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s)

	s.RegisterName("Arith", new(example.Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}

func addRegistryPlugin(s *server.Server) {
	clientConfig := constant.ClientConfig{
		TimeoutMs:            10 * 1000,
		ListenInterval:       30 * 1000,
		BeatInterval:         5 * 1000,
		NamespaceId:          "public",
		CacheDir:             "./cache",
		LogDir:               "./log",
		UpdateThreadNum:      20,
		NotLoadCacheAtStart:  true,
		UpdateCacheWhenEmpty: true,
	}

	serverConfig := []constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}}

	r := &nserverplugin.NacosRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		ClientConfig:   clientConfig,
		ServerConfig:   serverConfig,
		Cluster:        "test",
		Group:          "",
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
