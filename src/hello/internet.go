package hello

import (
	"fmt"
	"net"
)

//链路层（Mac），网络层（Ip），传输层（Tcp,端口),应用层（Interface+Parameter）

func Internet(){
	fmt.Println("this is hello.Internet")
	//监听
	listener,err1 := net.Listen("tcp","127.0.0.1:8080")
	if err1 != nil{
		fmt.Println("err = ",err1)
		return
	}

	defer listener.Close()
	//阻塞等待用户连接
	for  {
		listener.Accept()
	}

	//接收用户的请求


}