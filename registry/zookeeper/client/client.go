// go run -tags zookeeper client.go
package main

import (
	"context"
	"flag"
	"log"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/client"
)

var (
	zkAddr   = flag.String("zkAddr", "localhost:2181", "zookeeper address")
	basePath = flag.String("base", "/rpcx_test/Arith", "prefix path")
)

func main() {
	flag.Parse()

	d := client.NewZookeeperDiscovery(*basePath, []string{*zkAddr})
	xclient := client.NewXClient("Arith", "Mul", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}
	err := xclient.Call(context.Background(), args, reply, nil)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

}
