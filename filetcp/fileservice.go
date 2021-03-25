package main

import (
	"fmt"
	"net"
	"os"
	"io"
)

func main(){
	//创建服务用于监听
	listen,err1 := net.Listen("tcp","127.0.0.1:8884")
	if err1 != nil{
		fmt.Println("net.Listen err: ",err1)
		return
	}
	defer listen.Close()

	//阻塞等待用户连接，获取监听连接
	conn,err2 := listen.Accept()
	if err2 != nil{
		fmt.Println("listen.Accept err: ",err2)
		return
	}
	defer conn.Close()

	//获取数据,获取文件名，然后发送ok，最后再获取文件数据
	buf := make([]byte,1024)
	//获取文件名，先执行，成功后再发送ok给客户端
	var n int
	n,err1 = conn.Read(buf)
	if err1 != nil{
		fmt.Println("conn.Read err: ",err1)
		return
	}
	//这里说明已经可以获取到文件名，可以告诉客户端可以开始发送数据了，向客户端发送ok
	_,err3 := conn.Write([]byte("ok"))
	if err3 != nil {
		fmt.Println("conn.Write err: ",err3)
		return
	}
	//创建文件
	var filePath string
	filePath = "file_" + string(buf[:n])
	file,err3 := os.Create(filePath)
	if err3 != nil {
		fmt.Println("os.Create err: ",err3)
		return
	}
	defer file.Close()
	//开始往文件写内容
	SaveFile(conn,*file)

}

func SaveFile(conn net.Conn,file os.File){
	buf := make([]byte,1024)
	for {
		//获取监听到的数据
		n,err := conn.Read(buf)
		if err != nil {
			if err == io.EOF{
				fmt.Println("文件传输完毕")
			}else{
				fmt.Println("conn.Read err: ",err)	
			}
			return
		}
		//将数据保存到文件中
		file.Write(buf[:n])
	}
}