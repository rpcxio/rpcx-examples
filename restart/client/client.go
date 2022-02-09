package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/smallnest/rpcx/protocol"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/client"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	opt := client.DefaultOption
	opt.SerializeType = protocol.JSON

	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, opt)
	defer xclient.Close()

	args := example.Args{
		A: 10,
		B: 20,
	}

	// 第一次调用
	reply := &example.Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Printf("failed to call1: %v", err)
	}
	log.Printf("#1, %d * %d = %d", args.A, args.B, reply.C)

	time.Sleep(10 * time.Second) // 此时重启server

	// 第二次调用
	reply = &example.Reply{}
	err = xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Printf("failed to call2: %v", err)
	}
	log.Printf("#2, %d * %d = %d", args.A, args.B, reply.C)

	time.Sleep(10 * time.Second)

	// 第三次调用
	reply = &example.Reply{}
	err = xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Printf("failed to call3: %v", err)
	}
	log.Printf("#3, %d * %d = %d", args.A, args.B, reply.C)
}
