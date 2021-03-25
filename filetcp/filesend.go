package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//获取文件属性
	fmt.Println("请输入文件名称")
	var filePath string
	fmt.Scan(&filePath)

	info,err1 := os.Stat(filePath)		//获取文件属性，这里还未打开文件，需要要专门打开文件
	if err1 != nil {
		fmt.Println("os.Stat err: ",err1)
		return
	}

	//主动连接服务器
	conn,err2 := net.Dial("tcp","127.0.0.1:8884")
	if err2 != nil {
		fmt.Println("net.Dial err: ",err2)
		return
	}
	defer conn.Close()

	//给服务发送文件名
	_,err3 := conn.Write([]byte(info.Name()))
	if err3 != nil {
		fmt.Println("conn.Write err: ",err3)
		return
	}

	buf := make([]byte,1024)
	//接收服务器发送得ok，确认服务器已经做好接收准备
 	n,err4 := conn.Read(buf)
	if err4 != nil {
		fmt.Println("conn.Read Service err: ",err4)
		return
	}
	if "ok" == string(buf[:n]){
		//发送文件，先读文件，读多少发多少
		SendFile(filePath,conn)
	}
}

//发送文件
func SendFile(filePath string,conn net.Conn){
	//以只读的方式打开文件
	file,err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open err: ",err)
		return
	}
	defer file.Close()
	//暂存切片
	buf := make([]byte,1024*4)
	for {
		n,err1 := file.Read(buf)
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("文件发送完毕")
			}else{
				fmt.Println("file.Read err: ",err1)
			}
			return
		}
		conn.Write(buf[:n])
	}
}