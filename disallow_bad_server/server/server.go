package main

import (
	"context"
	"flag"
	"time"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
)

var (
	addr1 = flag.String("addr1", "localhost:8972", "server address")
	addr2 = flag.String("addr2", "localhost:8973", "server address")
)

type Arith int

func (t *Arith) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B
	return nil
}

type Arith2 int

func (t *Arith2) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B * 100
	return nil
}

func main() {
	flag.Parse()

	go func() {
		s := server.NewServer()
		s.RegisterName("Arith", new(Arith), "")
		s.Serve("reuseport", *addr1)
	}()

	go func() {
		s := server.NewServer()
		s.RegisterName("Arith", new(Arith2), "")
		go s.Serve("reuseport", *addr2)

		time.Sleep(time.Minute) // 去启动client
		s.Shutdown(context.Background())

		time.Sleep(5 * time.Minute)
		s = server.NewServer()
		s.RegisterName("Arith", new(Arith2), "")
		s.Serve("reuseport", *addr2)
	}()

	select {}
}
