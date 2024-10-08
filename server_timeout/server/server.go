package main

import (
	"context"
	"flag"
	"time"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

type Arith int

func (t *Arith) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	time.Sleep(200 * time.Second)
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer(
		server.WithReadTimeout(time.Duration(120)*time.Second),
		server.WithWriteTimeout(time.Duration(120)*time.Second),
	)
	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", *addr)
}
