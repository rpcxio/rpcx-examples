package main

import (
	"context"
	"flag"
	"time"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

var count = 0

type Arith int

func (t *Arith) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B
	count++
	if count%2 == 0 {
		time.Sleep(time.Minute)
	}
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	s.Register(new(Arith), "")
	s.Serve("tcp", *addr)
}
