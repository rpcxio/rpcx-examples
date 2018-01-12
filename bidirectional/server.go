package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"time"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

var clientConn net.Conn
var connected = false

type Arith int

func (t *Arith) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	clientConn = ctx.Value(server.RemoteConnContextKey).(net.Conn)
	reply.C = args.A * args.B
	connected = true
	return nil
}

func main() {
	flag.Parse()

	ln, _ := net.Listen("tcp", ":9981")
	go http.Serve(ln, nil)

	s := server.NewServer()
	//s.RegisterName("Arith", new(example.Arith), "")
	s.Register(new(Arith), "")
	go s.Serve("tcp", *addr)

	for !connected {
		time.Sleep(time.Second)
	}

	fmt.Printf("start to send messages to %s\n", clientConn.RemoteAddr().String())
	for {
		if clientConn != nil {
			err := s.SendMessage(clientConn, "test_service_path", "test_service_method", nil, []byte("abcde"))
			if err != nil {
				fmt.Printf("failed to send messsage to %s: %v\n", clientConn.RemoteAddr().String(), err)
				clientConn = nil
			}
		}
		time.Sleep(time.Second)
	}
}
