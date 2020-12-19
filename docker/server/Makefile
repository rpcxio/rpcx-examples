export GO= O111MODULE=on GOSUMDB=off GOPROXY=https://goproxy.cn,direct go


linux-bin:
	GOOS=linux $(GO) build -o server server.go 

image:linux-bin
	docker build -f ./Dockerfile -t smallnest/rpcx_server_demo:1.0 .

clean:
	rm -f server