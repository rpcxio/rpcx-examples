package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/smallnest/rpcx/serverplugin"

	"github.com/smallnest/rpcx/server"
)

var (
	addr             = flag.String("addr", "localhost:8972", "server address")
	fileTransferAddr = flag.String("transfer-addr", "localhost:8973", "data transfer address")
)

func main() {
	flag.Parse()

	s := server.NewServer()

	p := serverplugin.NewFileTransfer(*fileTransferAddr, nil, downloadFilehandler, 1000)
	serverplugin.RegisterFileTransfer(s, p)

	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}

func downloadFilehandler(conn net.Conn, args *serverplugin.DownloadFileArgs) {
	fmt.Printf("received file name: %s\n", args.FileName)

	f, err := os.Open("abc.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(conn, f)
	conn.Close()
	fmt.Println("send ok")
}
