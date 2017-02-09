package ClientServer

import "net"

func ServerProcedure()  {
	tcpAddr,err :=net.ResolveTCPAddr("tcp",":1200")
	checkError(err)
	listener,err:=net.ListenTCP("tcp",tcpAddr)
	checkError(err)

	for{
		connection,err:=listener.Accept()
		if err!=nil{
			continue
		}

		for{
			checkError(err)
			defer connection.Close()
			go handleConnection(connection)
		}
	}

}
