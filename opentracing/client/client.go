package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/kr/pretty"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/smallnest/rpcx/share"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	// create a mock tracer
	tracer := mocktracer.New()
	opentracing.SetGlobalTracer(tracer)

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	option := client.DefaultOption
	option.ReadTimeout = 10 * time.Second

	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	p := &client.OpenTracingPlugin{}
	pc := client.NewPluginContainer()
	pc.Add(p)
	xclient.SetPlugins(pc)

	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}
	ctx := context.WithValue(context.Background(), share.ReqMetaDataKey, map[string]string{"aaa": "from client"})
	ctx = context.WithValue(ctx, share.ResMetaDataKey, make(map[string]string))
	err := xclient.Call(ctx, "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
	log.Printf("received meta: %+v", ctx.Value(share.ResMetaDataKey))

	// print trace result
	spans := tracer.FinishedSpans()
	pretty.Print(spans)
}
