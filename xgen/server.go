package xgen

import (
	"context"
	"fmt"
	"time"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
)

type Arith int

func (t *Arith) Mul(ctx context.Context, args example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B
	return nil
}

func (t *Arith) Add(ctx context.Context, args *example.Args, reply *example.Reply) error {
	reply.C = args.A + args.B
	return nil
}

type Echo string

func (s *Echo) Echo(ctx context.Context, args string, reply *string) error {
	*reply = fmt.Sprintf("Hello %s from %s", args, *reply)
	return nil
}

type TimeS struct{}

func (s *TimeS) Time(ctx context.Context, args time.Time, reply *time.Time) error {
	*reply = time.Now()
	return nil
}
