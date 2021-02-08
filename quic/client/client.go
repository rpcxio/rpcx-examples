//go run -tags quic client.go
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "server address")
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

func main() {
	flag.Parse()

	// CA
	caCertPEM, err := ioutil.ReadFile("../ca.pem")
	if err != nil {
		panic(err)
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(caCertPEM)
	if !ok {
		panic("failed to parse root certificate")
	}

	conf := &tls.Config{
		// InsecureSkipVerify: true,
		RootCAs: roots,
	}

	option := client.DefaultOption
	option.TLSConfig = conf

	d, _ := client.NewPeer2PeerDiscovery("quic@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	args := &Args{
		A: 10,
		B: 20,
	}

	start := time.Now()
	for i := 0; i < 100000; i++ {
		reply := &Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
	}
	t := time.Since(start).Nanoseconds() / int64(time.Millisecond)

	fmt.Printf("tps: %d calls/s\n", 100000*1000/int(t))
}
