package main

import (
	"flag"
	"fmt"
	"time"

	metrics "github.com/rcrowley/go-metrics"
	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr     = flag.String("addr", "localhost:0", "server address")
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd address")
	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()

	s := server.NewServer()

	go s.Serve("tcp", *addr)

	time.Sleep(2 * time.Second)
	*addr = s.Address().String()
	fmt.Println("listened on: ", *addr)

	r := createRegistryPlugin()
	s.Plugins.Add(r)
	s.RegisterName("Arith", new(example.Arith), "")
	r.Start()

	select {}
}
func createRegistryPlugin() *serverplugin.EtcdRegisterPlugin {

	r := &serverplugin.EtcdRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		EtcdServers:    []string{*etcdAddr},
		BasePath:       *basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}

	return r
}
