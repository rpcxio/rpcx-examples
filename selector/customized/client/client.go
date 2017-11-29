package main

import (
	"context"
	"flag"
	"log"
	"sort"
	"strings"
	"time"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/client"
)

var (
	addr1 = flag.String("addr1", "tcp@localhost:8972", "server address")
	addr2 = flag.String("addr2", "tcp@localhost:8973", "server address")
)

func main() {
	flag.Parse()

	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1}, {Key: *addr2}})
	xclient := client.NewXClient("Arith", client.Failtry, client.SelectByUser, d, client.DefaultOption)
	defer xclient.Close()

	xclient.SetSelector(&alwaysFirstSelector{})

	args := &example.Args{
		A: 10,
		B: 20,
	}

	for i := 0; i < 10; i++ {
		reply := &example.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)

		time.Sleep(time.Second)
	}

}

type alwaysFirstSelector struct {
	servers []string
}

func (s *alwaysFirstSelector) Select(ctx context.Context, servicePath, serviceMethod string, args interface{}) string {
	var ss = s.servers
	if len(ss) == 0 {
		return ""
	}

	return ss[0]
}

func (s *alwaysFirstSelector) UpdateServer(servers map[string]string) {
	var ss = make([]string, 0, len(servers))
	for k := range servers {
		ss = append(ss, k)
	}

	sort.Slice(ss, func(i, j int) bool {
		return strings.Compare(ss[i], ss[j]) <= 0
	})
	s.servers = ss
}
