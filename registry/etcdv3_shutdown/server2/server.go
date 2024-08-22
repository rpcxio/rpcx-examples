package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/rpcxio/rpcx-etcd/serverplugin"
	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
)

var (
	addr     = flag.String("addr", "localhost:8973", "server address")
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd address")
	basePath = flag.String("base", "/rpcx_test", "prefix path")
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
				// panic(err)
			}
		}()

		time.Sleep((10 * time.Second))
		log.Println("server stopping")
		s.Shutdown(context.Background())
		log.Println("server stopped")
		time.Sleep((10 * time.Second))
	}

}

func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		EtcdServers:    []string{*etcdAddr},
		BasePath:       *basePath,
		UpdateInterval: time.Second,
		Expired:        time.Second * 2,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
