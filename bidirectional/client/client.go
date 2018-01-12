package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	c := client.NewClient(client.DefaultOption)
	err := c.Connect("tcp", *addr)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	// invoke once and set up the connection
	reply := &example.Reply{}
	err = c.Call(context.Background(), "Arith", "Mul", args, reply)

	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

	ch := make(chan *protocol.Message)
	c.RegisterServerMessageChan(ch)

	for msg := range ch {
		fmt.Printf("receive msg from server: %s\n", msg.Payload)
	}
}
