package main

import(
	"net"
	"bufio"
	"os"
	"fmt"
	"time"
)

func main(){
	
	service := ":1202"
	tcpAddr,err:=net.ResolveTCPAddr("tcp",service)
	checkErrorForServer(err)

	listener,err := net.ListenTCP("tcp",tcpAddr)
	checkErrorForServer(err)
	

	for{
		conn,err :=listener.Accept()
		if err!=nil{
			continue
		}
		fmt.Println("Connection @ Server: ",conn.RemoteAddr(),conn.LocalAddr())


		for{
			
			done :=make(chan bool,1)
			checkErrorForServer(err)
			go handleConnForServer(done,conn)
			<-done
						
		}
	}
}


func handleConnForServer(done chan bool, conn net.Conn){

	var buf [512]byte
	reader := bufio.NewReader(os.Stdin)
	for{
		n, err := conn.Read(buf[0:])
		checkErrorForServer(err)
		fmt.Println("Client : ", string(buf[:n]))
		time.Sleep(time.Second)
		fmt.Println("Enter text  : ")
		text, err := reader.ReadString('\n')
		checkErrorForServer(err)
		conn.Write([]byte(text))
		checkErrorForServer(err)
	}
	done<-true
}

func checkErrorForServer(err error){
	if err!=nil{
		fmt.Fprintf(os.Stderr,"Fatal error %s ",err.Error());
		os.Exit(1)
	}
}
