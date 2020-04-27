package main

//it is the general model about TCP server: the code from @studygolang

//to handle connection events
func HandleConn(conn net.Conn) {
	defer conn.Close()

	for {
		//read something from conn...
		//write something to conn...
	}
}

func main() {
	//listen to port 8888
	listen, err := net.Listen("tcp", ":8888")
	if nil != err {
		fmt.println("Listen error: ", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if nil != err {
			fmt.Println("Accept error: ", err)
			break
		}

		//start a new goroutine to handle the new connection
		go HandleConn(conn)
	}
}
