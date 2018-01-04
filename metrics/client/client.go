package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"time"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	go func() {
		for {
			reply := &example.Reply{}
			err := xclient.Call(context.Background(), "Mul", args, reply)
			if err != nil {
				log.Fatalf("failed to call: %v", err)
			}

			//log.Printf("%d * %d = %d", args.A, args.B, reply.C)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}

	}()

	go func() {
		for {
			reply := &example.Reply{}
			err := xclient.Call(context.Background(), "Add", args, reply)
			if err != nil {
				log.Fatalf("failed to call: %v", err)
			}

			//log.Printf("%d + %d = %d", args.A, args.B, reply.C)
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		}

	}()

	select {}
}
