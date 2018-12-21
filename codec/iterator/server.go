package main

import (
	"flag"

	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/share"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/rpcx-ecosystem/rpcx-examples3/codec/iterator/codec"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	share.Codecs[protocol.SerializeType(4)] = &codec.JsoniterCodec{}
	s := server.NewServer()
	//s.RegisterName("Arith", new(example.Arith), "")
	s.Register(new(example.Arith), "")
	s.Serve("tcp", *addr)
}
