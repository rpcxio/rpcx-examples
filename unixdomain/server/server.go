package main

import (
	"context"
	"flag"
	"os"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/v6/server"
)

var (
	addr = flag.String("addr", "/tmp/rpcx.socket", "server address")
)

type Arith struct{}

// the second parameter is not a pointer
func (t *Arith) Mul(ctx context.Context, args example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()

	os.Remove(*addr)
	s := server.NewServer()
	s.Register(new(Arith), "")
	s.Serve("unix", *addr)

}
