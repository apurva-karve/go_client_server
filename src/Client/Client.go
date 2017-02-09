package Client

import (
	"net/rpc"
	"log"
	"server"
	"fmt"
)

func ConnectToServer(){
	client, err := rpc.DialHTTP("tcp", "localhost" + ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	fmt.Println("Client is runnig...")
	args := &server.Args{7,8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}
