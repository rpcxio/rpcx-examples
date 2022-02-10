package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/nacos-group/nacos-sdk-go/common/constant"
	example "github.com/rpcxio/rpcx-examples"
	nserverplugin "github.com/rpcxio/rpcx-nacos/serverplugin"
	"github.com/smallnest/rpcx/server"
)

var addr = flag.String("addr", ":8972", "server address")

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

	host := os.Getenv("Nacos_Host")
	port := os.Getenv("Nacos_Port")
	nacosPort, _ := strconv.Atoi(port)

	serverConfig := []constant.ServerConfig{{
		IpAddr: host,
		Port:   uint64(nacosPort),
	}}

	addr := os.Getenv("MY_POD_IP") + ":8972"

	fmt.Println("register:", addr)

	r := &nserverplugin.NacosRegisterPlugin{
		ServiceAddress: "tcp@" + addr,
		ClientConfig:   clientConfig,
		ServerConfig:   serverConfig,
		Cluster:        "test",
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
