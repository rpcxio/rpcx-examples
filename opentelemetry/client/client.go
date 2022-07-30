package main

import (
	"context"
	"flag"
	"log"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/share"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"

	example "github.com/rpcxio/rpcx-examples"
	rotel "github.com/rpcxio/rpcx-plugins/client/otel"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	flag.Parse()

	tp := setOpenTelemetry()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	option := client.DefaultOption

	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	tracer := otel.Tracer("rpcx")
	p := rotel.NewOpenTelemetryPlugin(tracer, nil)

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

	if err := tp.Shutdown(context.Background()); err != nil {
		panic(err)
	}
}

func setOpenTelemetry() *trace.TracerProvider {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		panic(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithBatcher(exporter),
	)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetTracerProvider(tp)

	return tp
}
