package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/smallnest/rpcx/protocol"

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
	opt.SerializeType = protocol.JSON

	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, opt)
	defer xclient.Close()

	args := example.Args{
		A: 10,
		B: 20,
	}

	for {
		err := xclient.Oneshot(context.Background(), "Mul", args)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("send the problem %d * %d to server", args.A, args.B)
		time.Sleep(time.Second)
	}

}
