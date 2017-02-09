package main

import (
	"net"
	"fmt"
	"os"
	"time"
	"strings"
	"strconv"
)

func main(){

	service := ":"+os.Args[1]
	tcpAddr,_:=net.ResolveTCPAddr("tcp",service)
	listener,_ := net.ListenTCP("tcp",tcpAddr)
	addressValueMap := make(map[int]int)

	rangeStart, _ := strconv.Atoi(os.Args[2])
	rangeEnd, _ := strconv.Atoi(os.Args[3])
	for i:= rangeStart; i< rangeEnd; i++{
		addressValueMap[i]= -1
	}


	for{
		conn,err :=listener.Accept()
		if err!=nil{
			continue
		}
		fmt.Println("Connection @ Server: ",conn.RemoteAddr(),conn.LocalAddr())

		for{
			done :=make(chan bool,1)
			go readWrite(addressValueMap, done,conn)
			<-done
			fmt.Print("<< ENDS >>")
		}
	}
}

func readWrite(addressValueMap map[int]int, done chan bool, conn net.Conn){

	var buf [512]byte
	//reader := bufio.NewReader(os.Stdin)
	for{
		n, _ := conn.Read(buf[0:])
		msgSent := string(buf[:n])
		fmt.Println("Client : ", msgSent)

		time.Sleep(time.Second)

		var splitMsg = strings.Split(msgSent, " ")
		var address, _  = strconv.Atoi(splitMsg[1])

		if(!isAddressAvailable(address)){
			fmt.Fprintf(os.Stdout,"Not available here\n")
			secondServerAddress :=os.Args[2]
			tcpAddr,_ := net.ResolveTCPAddr("tcp4", secondServerAddress)
			conn,_:=net.DialTCP("tcp",nil,tcpAddr)
			fmt.Println("Connected to the other server")
			conn.Write([]byte(msgSent))
			fmt.Println("Forwarded command to the other server")
			n, _ := conn.Read(buf[0:])
			fmt.Println("Client : ", string(buf[:n]))
		}else{
			fmt.Fprintf(os.Stdout,"Available here\n")
			command :=splitMsg[0]
			if(command == "WRITE"){
				value,_ :=strconv.Atoi(splitMsg[2])
				writeToAddress(addressValueMap, address, value)
				conn.Write([]byte("WRITTEN"))
			}
			if(command == "READ"){
				val :=string(readFromAddress(addressValueMap, address))
				fmt.Fprintf(os.Stdout,"READ")
				fmt.Println("Reading....")
				conn.Write([]byte(val))
			}

		}
	}
	done <- true
}
func writeToAddress(addrmap map[int]int, address int, value int)bool {
	addrmap[address]=value
	return true
}
func readFromAddress(addrmap map[int]int, address int)int {
	return addrmap[address]

}

func isAddressAvailable(address int) bool{
	//if value, ok :=addressMap[address]; ok {
	//	fmt.Println(value)
	//	return true
	//}else{
	//	fmt.Println(value)
	//	return false
	//}
	rangeStart, _ := strconv.Atoi(os.Args[3])
	rangeEnd, _ := strconv.Atoi(os.Args[4])
	//fmt.Println(rangeStart, rangeEnd)
	if(address>rangeStart && address<rangeEnd){
		return true
	}
	return false
}
