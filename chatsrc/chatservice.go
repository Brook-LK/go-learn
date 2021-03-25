package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("chat.service")
	//监听
	listener,err1 := net.Listen("tcp","127.0.0.1:8888")	//network为tcp还是udp，address为地址端口（本地地址可以不用写:8888）
	if(err1 !=nil){
		fmt.Println("err1 = ",err1)
		return
	}
	defer listener.Close()

	//阻塞等待用户链接
	for {
		conn,err2 := listener.Accept()
		defer conn.Close()
		if(err2 !=nil){
			fmt.Println("err2 = ",err2)
			continue
		}
		//接收用户的请求
		buf := make([]byte,1024)
		n,err3 := conn.Read(buf)
		if(err3 != nil){
			fmt.Println("err3 = ",err3)
			continue
		}
		fmt.Println("buf = ",string(buf[:n]))

	}

}
