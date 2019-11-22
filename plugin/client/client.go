package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

var (
	xclient client.XClient
)

func main() {
	flag.Parse()

	go http.ListenAndServe(":9099", nil)
	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient = client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	plugins := client.NewPluginContainer()
	plugins.Add(&ConnectionPlugin{})
	xclient.SetPlugins(plugins)

	args := example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

}

type ConnectionPlugin struct {
}

func (p *ConnectionPlugin) ClientConnected(conn net.Conn) (net.Conn, error) {
	log.Printf("server %v connected", conn.RemoteAddr().String())

	args := example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

	return conn, nil
}

func (p *ConnectionPlugin) ClientConnectionClose(conn net.Conn) error {
	log.Printf("server %v closed", conn.RemoteAddr().String())
	return nil
}
