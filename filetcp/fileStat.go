package main

import (
	"fmt"
	"os"
)

func main() {
	//获取文件属性
	list := os.Args
	if len(list) != 2 {
		fmt.Println("useage: XXX file")
		return
	}

	fileName := list[1]

	info,err1 := os.Stat(fileName)		//获取文件属性
	if err1 != nil {
		fmt.Println("os.Stat err",err1)
		return
	}
	fmt.Println(info.Name())		//获取文件名
	fmt.Println(info.Size())		//获取文件大小
	fmt.Println(info.Mode())		//获取文件权限
	fmt.Println(info.ModTime())
	fmt.Println(info.IsDir())		//获取目录
	fmt.Println(info.Sys())
}
