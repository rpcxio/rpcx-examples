package main

import (
	"context"
	"flag"
	"log"
	"net"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
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
	s.Plugins.Add(&ConnectionPlugin{})

	//s.Register(new(Arith), "")
	s.RegisterName("Arith", new(Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}

type ConnectionPlugin struct {
}

func (p *ConnectionPlugin) HandleConnAccept(conn net.Conn) (net.Conn, bool) {
	log.Printf("client %v connected", conn.RemoteAddr().String())
	return conn, true
}

func (p *ConnectionPlugin) HandleConnClose(conn net.Conn) bool {
	log.Printf("client %v closed", conn.RemoteAddr().String())
	return true
}
