package main

import (
	"context"
	"flag"
	"log"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	oneClient := client.NewOneClient(client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer oneClient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}
	err := oneClient.Call(context.Background(), "Arith", "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

	var echo string
	err = oneClient.Call(context.Background(), "echo", "Say", "hello world", &echo)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	log.Printf("echo %s", echo)

}
