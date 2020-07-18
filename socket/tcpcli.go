package main

import (
	"fmt"
	"net"
	"time"
)

func ReadMsg(conn net.Conn){
	defer conn.Close()
	var buf = make([]byte, 64)
	for {
		n, err := conn.Read(buf)
		if nil != err {
			fmt.Println("Read error")
			break
		}
		fmt.Printf("read msg(%dB): %s\n", n, buf)
	}
}

func main(){
	conn, err := net.Dial("tcp","10.2.3.141:6666")
	if nil != err {
		fmt.Println("Connect to server error: ", err)
		return 
	}
	for i := 0; i < 100; i++{
		_, err := conn.Write([]byte("Hello world"))
		if nil != err {
			fmt.Println("Send message error: ", err)
			break
		}
		go ReadMsg(conn)		
		time.Sleep(time.Second*1)
	}
}
