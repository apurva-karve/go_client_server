package main

import(
	"fmt"
	"os"
	"net"
	"bufio"
)
func main(){

	if len(os.Args)!=2{
		fmt.Fprintf(os.Stderr,"Usage : %s host:port",os.Args[0])
		os.Exit(1)
	}
	service :=os.Args[1]

	tcpAddr,err := net.ResolveTCPAddr("tcp4",service)
	checkErrorForClient(err)

	conn,err:=net.DialTCP("tcp",nil,tcpAddr)
	//fmt.Println("Connection @ client: ",conn.RemoteAddr(),conn.LocalAddr())
	checkErrorForClient(err)
	//conn.SetTimeout(10000)
	//checkError(err)

	for{
		fmt.Fprintf(os.Stderr, "1")
		handleConnForClient(conn)
	}
	os.Exit(0)
}


func handleConnForClient(conn net.Conn){
	//var buf [512]byte
	reader := bufio.NewReader(os.Stdin)
		
	for{
		fmt.Println("Enter text  : ")
		text,err :=reader.ReadString('\n')
		checkErrorForClient(err)
		_,err =conn.Write([]byte(text))
		checkErrorForClient(err)
	}
}

func checkErrorForClient(err error){
	if err!=nil{
		fmt.Fprintf(os.Stderr,"Fatal error %s ",err.Error());
		os.Exit(1)
	}
}


