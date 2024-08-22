package main

import (
	"context"
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
	addr       = flag.String("addr", "localhost:8973", "server address")
	consulAddr = flag.String("consulAddr", "localhost:8500", "consul address")
	basePath   = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()

	for {
		s := server.NewServer()
		addRegistryPlugin(s)

		s.RegisterName("Arith", new(example.Arith), "")

		go func() {
			err := s.Serve("tcp", *addr)
			if err != nil {
				fmt.Println(err)
			}
		}()

		time.Sleep(time.Second * 10)
		log.Println("server stopping")
		s.Shutdown(context.Background())
		log.Println("server stopped")
		time.Sleep(time.Second * 10)
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
