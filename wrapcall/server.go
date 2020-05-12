package main

import (
	"context"
	"flag"
	"log"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()

	p := &WrapCall{}
	s.Plugins.Add(p)

	s.RegisterName("Greeter", new(example.Greeter), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}

type WrapCall struct{}

func (p *WrapCall) PreCall(ctx context.Context, serviceName, methodName string, args interface{}) (interface{}, error) {
	log.Printf("before %s.%s: args: %v", serviceName, methodName, args)
	if serviceName == "Greeter" && methodName == "Say" {
		name := args.(*string)
		if *name == "jack" {
			newName := "rose"
			return &newName, nil
		}
	}

	return args, nil
}

func (p *WrapCall) PostCall(ctx context.Context, serviceName, methodName string, args, reply interface{}) (interface{}, error) {
	log.Printf("after %s.%s: args: %v", serviceName, methodName, args)
	if serviceName == "Greeter" && methodName == "Say" {
		name := reply.(*string)
		if *name == "hello rose!" {
			newReply := "hello rose!!!"
			return &newReply, nil
		}
	}

	return reply, nil
}
