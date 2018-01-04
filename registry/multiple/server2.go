package main

// import (
// 	"context"
// 	"flag"

// 	example "github.com/rpcx-ecosystem/rpcx-examples3"
// 	"github.com/smallnest/rpcx/server"
// )

// var (
// 	addr2 = flag.String("addr2", "localhost:9981", "server2 address")
// )

// type Arith2 int

// func (t *Arith2) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
// 	reply.C = args.A * args.B * 100
// 	return nil
// }
// func main() {
// 	flag.Parse()

// 	s := server.NewServer()
// 	s.RegisterName("Arith", new(Arith2), "")
// 	s.Serve("tcp", *addr2)

// 	select {}
// }
