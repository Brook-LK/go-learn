package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("chat.client")
	//主动连接服务器
	conn,err1 := net.Dial("tcp","127.0.0.1:8888")
	defer conn.Close()
	if err1 != nil{
		fmt.Println("err1 = ",err1)
		return
	}

	//发送数据
	conn.Write([]byte("are you ok?"))


}
