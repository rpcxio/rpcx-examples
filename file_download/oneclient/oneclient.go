package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	oneClient := client.NewOneClient(client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer oneClient.Close()

	f, err := os.Create("abc.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = oneClient.DownloadFile(context.Background(), "abc.txt", f, map[string]string{"foo": "bar"})
	if err != nil {
		panic(err)
	}
	log.Println("received")

}
