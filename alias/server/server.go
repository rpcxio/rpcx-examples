package main

import (
	"flag"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/v6/server"
	"github.com/smallnest/rpcx/v6/serverplugin"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	a := serverplugin.NewAliasPlugin()
	a.Alias("a.b.c.D", "Times", "Arith", "Mul")
	s := server.NewServer()
	s.Plugins.Add(a)
	s.RegisterName("Arith", new(example.Arith), "")
	err := s.Serve("reuseport", *addr)
	if err != nil {
		panic(err)
	}
}
