package main

import (
	"context"
	"flag"
	"log"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	opt := client.DefaultOption
	opt.Heartbeat = false
	opt.TCPKeepAlivePeriod = 0
	opt.Retries = 0

	xclient := client.NewXClient("Arith", client.Failfast, client.RandomSelect, d, opt)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	log.Println("start to call")
	reply := &example.Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

}
