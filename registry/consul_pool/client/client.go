package main

import (
	"context"
	"flag"
	"log"
	"sync"
	"time"

	cclient "github.com/rpcxio/rpcx-consul/client"
	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/client"
)

var (
	consulAddr = flag.String("consulAddr", "localhost:8500", "consul address")
	basePath   = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()

	clientPool := sync.Pool{New: func() interface{} {
		d, _ := cclient.NewConsulDiscovery(*basePath, "Arith", []string{*consulAddr}, nil)
		xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
		return xclient
	}}

	args := &example.Args{
		A: 10,
		B: 20,
	}

	for {
		xclient := clientPool.Get().(client.XClient)
		reply := &example.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Printf("ERROR failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(1e9)
		clientPool.Put(xclient)
	}

}
