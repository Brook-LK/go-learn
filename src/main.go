package main

import (
	"fmt"
	"hello"
)

//main函数和init函数,当导入包中有init函数，会先执行导入包中的init函数，然后才执行本方法中的main函数
//如果本函数有init函数回在导入包init函数之后和main函数之前执行本包的init函数
//pkg文件夹是build的时候生成的，也可以由go install生成，先配置GOBIN
//src:放源代码，配置GOPATH
//bin：放可执行程序
//pkg：放需要的库
func init() {
	fmt.Println("this is main.main.init")
}

func main(){
	fmt.Println("hello")
	a := 1
	b := 10
	fmt.Println(hello.FmtReturn4(a,b))
	//hello.Main1()
	//hello.Compound()		//复合类型
	//hello.GuessNum()		//猜测随机数
	//hello.Lucky()
	//hello.Object()			//面向对象编程
	//hello.Interface()		//接口
	//hello.Err()				//异常错误处理
	//hello.StrUtil()			//字符串处理，正则，正则，文件处理
	//hello.Goroutine()		//并发
	//hello.Timmer()			//Timer定时
	//hello.Ticker()			//ticker定时
	hello.Internet()			//网络编程，这一部分单独拆出来写
}




