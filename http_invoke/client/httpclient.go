package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/smallnest/rpcx/codec"
)

const (
	XVersion           = "X-RPCX-Version"
	XMessageType       = "X-RPCX-MesssageType"
	XHeartbeat         = "X-RPCX-Heartbeat"
	XOneway            = "X-RPCX-Oneway"
	XMessageStatusType = "X-RPCX-MessageStatusType"
	XSerializeType     = "X-RPCX-SerializeType"
	XMessageID         = "X-RPCX-MessageID"
	XServicePath       = "X-RPCX-ServicePath"
	XServiceMethod     = "X-RPCX-ServiceMethod"
	XMeta              = "X-RPCX-Meta"
	XErrorMessage      = "X-RPCX-ErrorMessage"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

func main() {
	cc := &codec.MsgpackCodec{}

	args := &Args{
		A: 10,
		B: 20,
	}

	data, _ := cc.Encode(args)

	req, err := http.NewRequest("POST", "http://127.0.0.1:8972/", bytes.NewReader(data))
	if err != nil {
		log.Fatal("failed to create request: ", err)
		return
	}

	h := req.Header
	h.Set(XMessageID, "10000")
	h.Set(XMessageType, "0")
	h.Set(XSerializeType, "3")
	h.Set(XServicePath, "Arith")
	h.Set(XServiceMethod, "Mul")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("failed to call: ", err)
	}
	defer res.Body.Close()

	// handle http response
	replyData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("failed to read response: ", err)
	}

	reply := &Reply{}
	err = cc.Decode(replyData, reply)
	if err != nil {
		log.Fatal("failed to decode reply: ", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
