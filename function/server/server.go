package main

import (
	"context"
	"flag"

	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
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
	s.RegisterFunction("a.fake.service", mul, "")
	s.Serve("tcp", *addr)
}
