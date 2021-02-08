package main

import (
	"context"
	"flag"
	"log"
	"time"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/v6/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
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
