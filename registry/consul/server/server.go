package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	metrics "github.com/rcrowley/go-metrics"
	cserver "github.com/rpcxio/rpcx-consul/serverplugin"
	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
)

var (
	addr       = flag.String("addr", "localhost:8972", "server address")
	consulAddr = flag.String("consulAddr", "localhost:8500", "consul address")
	basePath   = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s)

	s.RegisterName("Arith", new(example.Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		fmt.Println(err)
	}
}

func addRegistryPlugin(s *server.Server) {

	r := &cserver.ConsulRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		ConsulServers:  []string{*consulAddr},
		BasePath:       *basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
