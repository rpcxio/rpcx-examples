package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"

	"github.com/smallnest/rpcx/v6/share"

	"github.com/smallnest/rpcx/v6/server"
)

var (
	addr             = flag.String("addr", "localhost:8972", "server address")
	fileTransferAddr = flag.String("transfer-addr", "localhost:8973", "data transfer address")
)

func main() {
	flag.Parse()

	s := server.NewServer()

	p := server.NewFileTransfer(*fileTransferAddr, saveFilehandler, nil, 1000)
	s.EnableFileTransfer(share.SendFileServiceName, p)

	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}

func saveFilehandler(conn net.Conn, args *share.FileTransferArgs) {
	fmt.Printf("received file name: %s, size: %d\n", args.FileName, args.FileSize)
	data, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Printf("error read: %v\n", err)
		return
	}
	fmt.Printf("file content: %s\n", data)
}
