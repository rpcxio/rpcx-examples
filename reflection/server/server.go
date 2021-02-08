package main

import (
	"flag"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/v6/reflection"
	"github.com/smallnest/rpcx/v6/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()

	p := reflection.New()
	s.Plugins.Add(p)

	s.Register(new(example.Arith), "")
	s.Register(p, "")
	s.Serve("tcp", *addr)
}
