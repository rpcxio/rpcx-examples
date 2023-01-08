package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
)

var (
	addr1 = flag.String("addr1", "localhost:8972", "server address")
	addr2 = flag.String("addr2", "localhost:8973", "server address")
)

type Arith struct{}

// the second parameter is not a pointer
func (t *Arith) Mul(ctx context.Context, args example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B * 10
	return nil
}

type Arith2 struct{}

// the second parameter is not a pointer
func (t *Arith2) Mul(ctx context.Context, args example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B * 20
	return nil
}

func main() {
	flag.Parse()

	// server1 is still runing
	s1 := server.NewServer()
	//s.Register(new(Arith), "")
	s1.RegisterName("Arith", new(Arith), "")
	go s1.Serve("tcp", *addr1)

	s2 := server.NewServer()
	//s.Register(new(Arith), "")
	s2.RegisterName("Arith", new(Arith2), "")
	go s2.Serve("tcp", *addr2)

	time.Sleep(time.Minute)
	err := s2.Shutdown(context.Background())
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Minute)
	fmt.Println("restart ...")
	s2 = server.NewServer()
	//s.Register(new(Arith), "")
	s2.RegisterName("Arith", new(Arith2), "")
	s2.Serve("tcp", *addr2)

}
