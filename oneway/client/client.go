package main

import (
	"context"
	"flag"
	"log"
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

	args := example.Args{
		A: 10,
		B: 20,
	}

	for i := 0; i < 10; i++ {
		err := xclient.Call(context.Background(), "Mul", args, nil)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}
		time.Sleep(time.Second)
	}

}
