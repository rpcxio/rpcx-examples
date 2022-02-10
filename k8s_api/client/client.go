package main

import (
	"context"
	"fmt"
	"log"
	"sort"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/client"
)

var lastIPs []string

func main() {
	clientset, pairs := getPodIPs("rpcx-server-demo2", 8972)
	d, _ := client.NewMultipleServersDiscovery(pairs)
	go watchPodIPs("rpcx-server-demo2", 8972, clientset, d)

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
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(1e9)
	}
}

func watchPodIPs(serverPod string, port int, clientset *kubernetes.Clientset, d *client.MultipleServersDiscovery) {
	tick := time.NewTicker(10 * time.Second)
	for range tick.C {

		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{
			LabelSelector: "app=rpcx-server-demo2",
		})
		if err != nil {
			continue
		}

		var ips []string
		for _, pod := range pods.Items {
			ips = append(ips, pod.Status.PodIP)
		}

		updateDiscovery(ips, 8972, d)
	}
}

func getPodIPs(serverPod string, port int) (*kubernetes.Clientset, []*client.KVPair) {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{
		LabelSelector: "app=rpcx-server-demo2",
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	var ips []string
	for _, pod := range pods.Items {
		ips = append(ips, pod.Status.PodIP)
	}

	var pairs []*client.KVPair
	for _, ip := range ips {
		pairs = append(pairs, &client.KVPair{
			Key: fmt.Sprintf("tcp@%s:%d", ip, port),
		})
	}

	return clientset, pairs
}

func updateDiscovery(ips []string, port int, d *client.MultipleServersDiscovery) {
	if len(ips) == 0 { // 如果没有可用的pod,我们认为是异常情况，保留上一次读取的pod
		return
	}

	sort.Strings(ips)
	if fmt.Sprintf("%v", ips) == fmt.Sprintf("%v", lastIPs) {
		return
	}
	lastIPs = ips

	var pairs []*client.KVPair
	for _, ip := range ips {
		pairs = append(pairs, &client.KVPair{
			Key: fmt.Sprintf("tcp@%s:%d", ip, port),
		})
	}

	d.Update(pairs)
}
