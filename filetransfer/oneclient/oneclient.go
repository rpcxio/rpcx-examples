package main

import (
	"context"
	"flag"
	"log"

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

	err := oneClient.Sendfile(context.Background(), "abc.txt", 0)
	if err != nil {
		panic(err)
	}
	log.Println("send ok")

}
