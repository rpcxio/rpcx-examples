package main

import (
	"context"
	"flag"
	"log"

	"github.com/smallnest/rpcx/protocol"

	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	opt := client.DefaultOption
	opt.SerializeType = protocol.JSON

	xclient := client.NewXClient("Greeter", client.Failtry, client.RandomSelect, d, opt)
	defer xclient.Close()

	var reply string
	err := xclient.Call(context.Background(), "Say", "jack", &reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%s", reply)

}
