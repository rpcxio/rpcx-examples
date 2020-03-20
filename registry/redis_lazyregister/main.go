package main

import (
	"context"
	"flag"
	"log"
	"time"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr      = flag.String("addr", "localhost:8972", "server address")
	redisAddr = flag.String("redisAddr", "localhost:6379", "redis address")
	basePath  = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()

	d := client.NewRedisDiscovery(*basePath, "Arith", []string{*redisAddr}, nil)
	xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}

	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Printf("call failed becuase the service has not registered: %v", err)
	}

	go startServer()

	time.Sleep(5 * time.Second)

	log.Printf("server started, begin to call again")

	for {
		err = xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("expect normal but got err: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(time.Second)
	}
}

func startServer() {
	s := server.NewServer()
	addRegistryPlugin(s)

	s.RegisterName("Arith", new(example.Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}
func addRegistryPlugin(s *server.Server) {

	r := &serverplugin.RedisRegisterPlugin{
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
