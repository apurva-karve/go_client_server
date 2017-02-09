package main

import(
	"fmt"
	"os"
	"net"
//	"io/ioutil"
	"bufio"
)
func main(){

	if len(os.Args)!=2{
		fmt.Fprintf(os.Stderr,"Usage : %s host:port",os.Args[0])
		os.Exit(1)
	}
	var buf [512]byte
	service :=os.Args[1]

	udpAddr,err := net.ResolveUDPAddr("udp4",service)
	checkError(err)

	conn,err:=net.DialUDP("udp",nil,udpAddr)

	checkError(err)

	/*n,err:=conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	fmt.Println("No. of bytes written are : ",n)
	checkError(err)*/

	reader :=bufio.NewReader(os.Stdin)
	for{
		fmt.Println("Enter text  : ")
		text,err :=reader.ReadString('\n')
		checkError(err)	
		n,err :=conn.Write([]byte(text))
		tp(n)
		nr,err:= conn.Read(buf[0:])			
				
		checkError(err)			
		fmt.Println("Server : ",string(buf[:nr]))
		/*result,err := ioutil.ReadAll(conn)
	
//		result,err := conn.Read()
		checkError(err)
	
		fmt.Println("Server : ",string(result))*/
	}
	//fmt.Printf("The type of result is  : %T \n",result)
	os.Exit(0)
}

func checkError(err error){
	if err!=nil{
		fmt.Fprintf(os.Stderr,"Fatal error %s ",err.Error());
		os.Exit(1)
	}
}

func tp(n int){
}
