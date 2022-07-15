package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	example "github.com/rpcxio/rpcx-examples"
	nclient "github.com/rpcxio/rpcx-nacos/client"
	"github.com/smallnest/rpcx/client"
)

func main() {
	flag.Parse()

	d, _ := configNacos()
	xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &example.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Printf("failed to call: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)

		time.Sleep(5 * time.Second)
	}
}

func configNacos() (client.ServiceDiscovery, error) {
	clientConfig := constant.ClientConfig{
		TimeoutMs:            10 * 1000,
		ListenInterval:       30 * 1000,
		BeatInterval:         5 * 1000,
		NamespaceId:          "public",
		CacheDir:             "./cache",
		LogDir:               "./log",
		UpdateThreadNum:      20,
		NotLoadCacheAtStart:  true,
		UpdateCacheWhenEmpty: true,
	}

	serverConfig := []constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}}

	return nclient.NewNacosDiscovery("Arith", "test", "", clientConfig, serverConfig)
}
