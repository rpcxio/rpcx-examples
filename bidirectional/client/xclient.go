package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	ch := make(chan *protocol.Message)

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewBidirectionalXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption, ch)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

	for msg := range ch {
		fmt.Printf("receive msg from server: %s\n", msg.Payload)
	}
}
