package main

import (
	"context"
	"flag"
	"log"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/share"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient(share.SendFileServiceName, client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	err := xclient.SendFile(context.Background(), "abc.txt", 0, map[string]string{"foo": "bar"})
	if err != nil {
		panic(err)
	}
	log.Println("send ok")

}
