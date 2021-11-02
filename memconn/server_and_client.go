// go run -tags kcp server.go
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
)

var addr = flag.String("addr", "a-test-server", "server address")

func main() {
	flag.Parse()

	s := server.NewServer()
	s.RegisterName("Arith", new(example.Arith), "")

	go func() {
		err := s.Serve("memu", *addr)
		if err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second)

	d, _ := client.NewPeer2PeerDiscovery("memu@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	start := time.Now()
	for i := 0; i < 10000; i++ {
		reply := &example.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}
		// log.Printf("%d * %d = %d", args.A, args.B, reply.C)
	}
	dur := time.Since(start)
	qps := 10000 * 1000 / int(dur/time.Millisecond)
	fmt.Printf("qps: %d call/s", qps)
}
