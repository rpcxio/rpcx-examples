//go run -tags quic client.go
package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"time"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	option := client.DefaultOption
	option.TLSConfig = conf

	d := client.NewPeer2PeerDiscovery("quic@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	start := time.Now()
	for i := 0; i < 100000; i++ {
		reply := &example.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
	}
	t := time.Since(start).Nanoseconds() / int64(time.Millisecond)

	fmt.Printf("tps: %d calls/s\n", 100000*1000/int(t))
}
