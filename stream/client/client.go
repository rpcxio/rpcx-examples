package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/share"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient(share.StreamServiceName, client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	// get a connection for streaming
	conn, err := xclient.Stream(context.Background(), map[string]string{"foo": "bar"})
	if err != nil {
		panic(err)
	}

	go func() {
		io.Copy(os.Stdout, conn)
		conn.Close()
	}()

	for i := 0; i < 10; i++ {
		fmt.Fprintf(conn, "#%d: %d\n", i, time.Now().Unix())
		time.Sleep(time.Second)
	}

	conn.Close()
}
