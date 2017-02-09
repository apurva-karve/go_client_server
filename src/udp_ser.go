package main

import(
	"fmt"
	"os"
	"net"
	//"time"
//	"io/ioutil"
	"bufio"
)

func main(){
	var buf [512]byte
	service := ":1200"
	udpAddr,err:=net.ResolveUDPAddr("udp4",service)
	checkError(err)

	//listener,err := net.ListenUDP("tcp",tcpAddr)
	conn,err :=net.ListenUDP("udp",udpAddr)	
	checkError(err)
	reader := bufio.NewReader(os.Stdin)

	for{
		//conn,err :=listener.Accept()
		if err!=nil{
			continue
		}
		/*daytime :=time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()*/
		
		for{
			n,addr,err:= conn.ReadFromUDP(buf[0:])
	//		result,err := conn.Read()
			
			checkError(err)			
			fmt.Println("Client : ",string(buf[:n]))	
			fmt.Println("Enter text  : ")
			text,err:=reader.ReadString('\n')
			checkError(err)	
			conn.WriteToUDP([]byte(text),addr)
			//tp(n)
			checkError(err)			
		}
	}
}

func checkError(err error){
	if err!=nil {
		fmt.Fprintf(os.Stderr,"Fatal error : %s ",err.Error());
		os.Exit(1)	
	}
}

/*func tp(n int){
}*/
