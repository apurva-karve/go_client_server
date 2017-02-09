package server

import (
	"net/rpc"
	"net"
	"log"
	"net/http"
	"fmt"
)

func RunServe() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	fmt.Println("Server is running....")
	go http.Serve(l, nil)
}