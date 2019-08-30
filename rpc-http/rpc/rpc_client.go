package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	A int
	B int
}

func main() {
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error in dialing.%s", err)
		return
	}
	var reply int
	var args = Args{A: 10, B: 20}
	err = client.Call("OperServer.Add", args, &reply)
	if err != nil {
		log.Fatalf("Error in Call.%s", err)
		return
	}
	log.Printf("%d + %d = %d", args.A, args.B, reply)
}
