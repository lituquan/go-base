package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type OperServer struct{}

//用结构体包装参数，适配client.Call 的参数
type Args struct {
	A int
	B int
}

func (*OperServer) Add(args Args, result *int) error {
	log.Println("a,b-->", args.A, args.B)
	*result = args.A + args.B
	return nil
}

var server = OperServer{}

func main() {
	rpc.Register(&server)
	rpc.HandleHTTP()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("server error", err)
	}
}
