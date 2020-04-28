package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/smallnest/rpcx/v5/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

type Arith struct{}

// the second parameter is not a pointer
func (t *Arith) Mul(ctx context.Context, args Args, reply *Reply) error {
	reply.C = args.A * args.B
	fmt.Println("C=", reply.C)
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	//s.Register(new(Arith), "")
	s.RegisterName("Arith", new(Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}
