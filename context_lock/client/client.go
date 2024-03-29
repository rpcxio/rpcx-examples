package main

import (
	"context"
	"flag"
	"log"
	"sync"

	"github.com/smallnest/rpcx/share"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")

	option := client.DefaultOption
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}
	ctx := context.WithValue(context.Background(), share.ReqMetaDataKey, map[string]string{"aaa": "from client"})
	ctx = context.WithValue(ctx, share.ResMetaDataKey, make(map[string]string))
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			args.A = i
			err := xclient.Call(ctx, "Mul", args, reply)
			if err != nil {
				log.Fatalf("failed to call: %v", err)
			}
			log.Printf("%d * %d = %d", args.A, args.B, reply.C)
			log.Printf("received meta: %+v", ctx.Value(share.ResMetaDataKey))
		}(i)

	}

	wg.Wait()
}
