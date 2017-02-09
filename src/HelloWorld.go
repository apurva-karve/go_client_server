package main

import "fmt"
import (
	"server"
	"Client"
)
func main(){
	fmt.Println("Hello World");
	server.RunServe()
	Client.ConnectToServer()
}