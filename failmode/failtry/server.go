package main

import (
	"context"
	"errors"
	"flag"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/server"
)

var (
	addr1 = flag.String("addr1", "localhost:8972", "server1 address")
	addr2 = flag.String("addr2", "localhost:9981", "server2 address")
)

type Arith int

func (t *Arith) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	*t = *t + 1
	if *t%2 == 0 {
		return errors.New("unknown error")
	}
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()

	go createServer(*addr1)
	go createServer(*addr2)

	select {}
}

func createServer(addr string) {
	s := server.NewServer()
	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", addr)
}
