package main

import (
	"flag"
	"log"
	"time"

	metrics "github.com/rcrowley/go-metrics"
	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
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
	s.Serve("tcp", *addr)
}

func addRegistryPlugin(s *server.Server) {

	r := &serverplugin.ConsulRegisterPlugin{
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
