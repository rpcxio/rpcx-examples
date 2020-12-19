package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

type Arith struct{}

// the second parameter is not a pointer
func (t *Arith) Mul(ctx context.Context, args example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	//s.Register(new(Arith), "")
	s.RegisterName("Arith", new(Arith), "")
	go func() {
		err := s.Serve("reuseport", *addr)
		if err != nil {
			panic(err)
		}
	}()

	// kill -HUP <server pid> : restart this server
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)
	<-c
	s.Restart(context.Background())
}
