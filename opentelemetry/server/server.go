package main

import (
	"context"
	"flag"
	"os"
	"io"

	"github.com/smallnest/rpcx/serverplugin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

type Arith int

func (t *Arith) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()

	tp := setOpenTelemetry()
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			panic(err)
		}
	}()


	s := server.NewServer()

	var tracer = otel.Tracer("rpcx")
	p := serverplugin.NewOpenTelemetryPlugin(tracer, nil)
	s.Plugins.Add(p)

	s.RegisterName("Arith", new(Arith), "")

	s.Serve("tcp", *addr)
}


func setOpenTelemetry() *trace.TracerProvider{
	exp, err := newExporter(os.Stdout)
	if err != nil {
		panic(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
	)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			panic(err)
		}
	}()
	otel.SetTracerProvider(tp)
	
	return tp

}


func newExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}