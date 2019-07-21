package main

import (
	"context"
	"flag"
	"log"

	"github.com/smallnest/rpcx/serverplugin"

	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient(serverplugin.SendFileServiceName, client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	err := xclient.SendFile(context.Background(), "abc.txt", 0)
	if err != nil {
		panic(err)
	}
	log.Println("send ok")

}
