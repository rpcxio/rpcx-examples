//go run -tags kcp client.go
package main

import (
	"context"
	"crypto/sha1"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/client"
	kcp "github.com/xtaci/kcp-go"
	"golang.org/x/crypto/pbkdf2"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

const cryptKey = "rpcx-key"
const cryptSalt = "rpcx-salt"

func main() {
	flag.Parse()

	pass := pbkdf2.Key([]byte(cryptKey), []byte(cryptSalt), 4096, 32, sha1.New)
	bc, _ := kcp.NewAESBlockCrypt(pass)
	option := client.DefaultOption
	option.Block = bc

	d := client.NewPeer2PeerDiscovery("kcp@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RoundRobin, d, option)
	defer xclient.Close()

	// plugin
	cs := &ConfigUDPSession{}
	pc := client.NewPluginContainer()
	pc.Add(cs)
	xclient.SetPlugins(pc)

	args := &example.Args{
		A: 10,
		B: 20,
	}

	start := time.Now()
	for i := 0; i < 10000; i++ {
		reply := &example.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}
		//log.Printf("%d * %d = %d", args.A, args.B, reply.C)
	}
	dur := time.Since(start)
	qps := 10000 * 1000 / int(dur/time.Millisecond)
	fmt.Printf("qps: %d call/s", qps)
}

type ConfigUDPSession struct{}

func (p *ConfigUDPSession) ConnCreated(conn net.Conn) (net.Conn, error) {
	session, ok := conn.(*kcp.UDPSession)
	if !ok {
		return conn, nil
	}

	session.SetACKNoDelay(true)
	session.SetStreamMode(true)
	return conn, nil
}
