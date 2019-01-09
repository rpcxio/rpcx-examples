package main

import (
	"context"
	"flag"
	"log"
	"os"

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

	f, err := os.Create("abc.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = xclient.DownloadFile(context.Background(), "abc.txt", f)
	if err != nil {
		panic(err)
	}
	log.Println("received")
}
