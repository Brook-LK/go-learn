package main

import (
	"fmt"
	"net"
	"strings"
)

//手写一个websocket
func main() {
	fmt.Println("chat.service")
	//监听
	listener,err1 := net.Listen("tcp","127.0.0.1:8882")	//network为tcp还是udp，address为地址端口（本地地址可以不用写:8888）
	if(err1 !=nil){
		fmt.Println("err1 = ",err1)
		return
	}
	defer listener.Close()
	//接收多个用户
	for {
		conn, err2 := listener.Accept()
		//defer conn.Close()
		if (err2 != nil) {
			fmt.Println("err2 = ", err2)
			continue
		}
		//处理用户请求，新建一个协程，一个用户占一个协程
		go HandleConn(conn)		//应该是一个用户用一个net.Conn,新增加用户新加一个协程

	}
}

func HandleConn(conn net.Conn)  {
	defer conn.Close()		//断开连接
	//获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Printf("客户端 %s 连接成功\n",addr)

	//读取用户发送的信息
	buf := make([]byte,1024)
	for {
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取客户发送信息失败",err)
			return
		}
		fmt.Printf("收到客户端 %s 发送的信息: %s\n",addr,string(buf[:n]))
		if string(buf[:n]) == "exit"{
			fmt.Println(addr," exit")
			return
		}

		//把数据转换为大写再发送给客户
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}


}