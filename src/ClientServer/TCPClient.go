package ClientServer

import (
	"net"
	"fmt"
	"os"
	"bufio"
)

func ClientProcedure()  {
	tcpAddress,err:=net.ResolveTCPAddr("tcp4","localhost:1200")
	checkError(err)
	connection,err:=net.DialTCP("tcp",nil,tcpAddress)
	checkError(err)

	for{
		defer connection.Close()
		handleConnection(connection)
	}

}
func handleConnection(connection net.Conn) {
	done :=make(chan bool)
	go writeToConnection(done,connection)
	<-done
	go readFromConnection(done,connection)
	<-done
}
func readFromConnection(done chan bool, conn net.Conn) {
	var buf [512]byte
	n,err:=conn.Read(buf[0:])
	checkError(err)
	if n!=0{
		fmt.Println("Other : ",string(buf[:n]))
	}
}
func writeToConnection(done chan bool, conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text : ")
	text,err:=reader.ReadString('\n')
	checkError(err)

	if len(text)>0{
		conn.Write([]byte(text))
	}
	done<-true
}
func checkError(err error) {
	if err!=nil{
		fmt.Fprintf(os.Stderr,"Fatal error %s",err.Error())
		os.Exit(1)
	}
}
