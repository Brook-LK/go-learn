package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	fmt.Println("chat.client")
	//主动连接服务器
	conn,err1 := net.Dial("tcp","127.0.0.1:8882")
	defer conn.Close()
	if err1 != nil{
		fmt.Println("err1 = ",err1)
		return
	}

	//发送数据
	//conn.Write([]byte("are you ok1?"))
	//开协程通过控制端输入发送消息
	go sendMess(conn)

	//读取从服务端发送过来的消息
	buf := make([]byte,1024)
	for {
		n,err2 := conn.Read(buf)
		if err2 != nil{
			if err2 == io.EOF{
				fmt.Println("即将断开连接")
				break
			}
			fmt.Println("获取服务端发送的数据失败",err2)
			continue
		}
		fmt.Println("客户端发送消息：",string(buf[:n]))
	}
}

func sendMess(conn net.Conn){
	for{
		var str string
		fmt.Scan(&str)
		conn.Write([]byte(str))
	}
}