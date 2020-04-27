package main 

import (
	"fmt"
	"net"
)

func HandleConn(conn net.Conn){
	defer conn.Close()
	var buf = make([]byte, 64)
	for{
		n, err := conn.Read(buf)
		if nil != err {
			fmt.Println("Read error: ",err)
			break
		} else {
			fmt.Printf("Recieve(%dB): %s\n", n, buf)
		}

		n, err = conn.Write([]byte("I have read your message"))
		if nil != err {
			fmt.Println("Write error: ",err)
			break
		} else {
			fmt.Printf("Write(%dB) success!\n",n)
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if nil != err {
		fmt.Println("Listen error: ",err)
		return
	}

	for {
		conn, err := listen.Accept()
		if nil != err {
			fmt.Println("Accept error:", err)
		}

		go HandleConn(conn)
	}
}
