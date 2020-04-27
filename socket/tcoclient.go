package main 

//the model for TCP client connecting to TCP server 
//1.Block Dial function: net.Dial()
//2.Timeout Dial function: net.DialTimeout()

func main() {
	//Block Dial function example
	conn, err := net.Dial("tcp", "10.2.3.141:8888")
	defer conn.Close()
	if nil != err {
		fmt.Println("Connect to TCP server error: ", err)
		return
	}
	
	//Timeout Dial function example
	/*
	conn, err := net.DialTimeout("tcp","10.2.3.141:8888", 2*time.Second)
	defer conn.Close()
	if nil != err {
		fmt.Println("Connect to TCP server tiomout")
		return
	}
	*/

	
}

