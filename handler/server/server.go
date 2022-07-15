package main

import (
	"flag"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func mul(ctx *server.Context) error {
	var args example.Args
	err := ctx.Bind(&args)
	if err != nil {
		return err
	}

	var reply example.Reply
	reply.C = args.A * args.B

	ctx.Write(reply)

	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()

	s.AddHandler("Arith", "Mul", mul)

	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}
