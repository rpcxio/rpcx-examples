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
	addr     = flag.String("addr", "localhost:8972", "server address")
	zkAddr   = flag.String("zkAddr", "localhost:2181", "zookeeper address")
	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s)

	s.RegisterName("Arith", new(example.Arith), "")
	s.Serve("tcp", *addr)
}

func addRegistryPlugin(s *server.Server) {

	r := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress:   "tcp@" + *addr,
		ZooKeeperServers: []string{*zkAddr},
		BasePath:         *basePath,
		Metrics:          metrics.NewRegistry(),
		UpdateInterval:   time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
