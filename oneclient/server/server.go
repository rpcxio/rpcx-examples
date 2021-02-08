package main

import (
	"context"
	"flag"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/v6/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

type Echo int

func (*Echo) Say(ctx context.Context, args string, reply *string) error {
	*reply = args
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	//s.RegisterName("Arith", new(example.Arith), "")
	s.Register(new(example.Arith), "")
	s.RegisterName("echo", new(Echo), "")
	s.Serve("tcp", *addr)
}
