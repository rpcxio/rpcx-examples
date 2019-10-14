package main

import (
	"context"
	"flag"
	"log"
	"time"

	metrics "github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr     = flag.String("addr", "localhost:8972", "server address")
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd address")
	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

func mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s)

	s.RegisterFunction("a.fake.service", mul, "")
	go s.Serve("tcp", *addr)

	time.Sleep(time.Minute)

	err := s.UnregisterAll()
	if err != nil {
		panic(err)
	}
}

func addRegistryPlugin(s *server.Server) {

	r := &serverplugin.EtcdRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		EtcdServers:    []string{*etcdAddr},
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
