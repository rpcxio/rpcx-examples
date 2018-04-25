package main

import (
	"context"
	"flag"
	"log"
	"strings"
	"time"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	option := client.DefaultOption
	option.CompressType = protocol.Gzip

	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	args := strings.Repeat("world", 2048)
	for {
		var reply string
		err := xclient.Call(context.Background(), "Say", args, &reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("replay = (len: %d) %s", len(reply), reply[:20])
		time.Sleep(1e9)
	}

}
