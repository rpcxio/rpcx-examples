package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/smallnest/rpcx/serverplugin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
)

var addr = flag.String("addr", "localhost:8972", "server address")

type Arith int

func (t *Arith) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()

	tp := setOpenTelemetry()

	s := server.NewServer()

	tracer := otel.Tracer("rpcx")
	p := serverplugin.NewOpenTelemetryPlugin(tracer, nil)
	s.Plugins.Add(p)
	s.RegisterName("Arith", new(Arith), "")

	go func() {
		s.Serve("tcp", *addr)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

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
