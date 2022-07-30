package main

import (
	"flag"
	"log"
	"time"

	example "github.com/rpcxio/rpcx-examples"
	cserver "github.com/rpcxio/rpcx-redis/serverplugin"
	"github.com/smallnest/rpcx/server"
)

var (
	addr      = flag.String("addr", "localhost:8972", "server address")
	redisAddr = flag.String("redisAddr", "localhost:6379", "redis address")
	basePath  = flag.String("base", "/rpcx_test", "prefix path")
)

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

	r := &cserver.RedisRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		RedisServers:   []string{*redisAddr},
		BasePath:       *basePath,
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
