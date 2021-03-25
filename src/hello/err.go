package hello

import(
	"fmt"
	"errors"
)

//error为普通错误，panic为致命错误无法恢复的，比如数组越界就是调用的panic，recover函数用于恢复panic错误导致停止的程序，必须放在defer后面
func Err(){
	fmt.Println("there is hello.Err")
	//错误接口1
	//var err1 error = fmt.Errorf("%s","this is normal err")
	err1 := fmt.Errorf("%s","this is normal err")
	fmt.Println("err1 = ",err1)
	//创建错误接口
	err2 := errors.New("this is new err")
	fmt.Println("err2 = ",err2)

	//error接口的应用，用一个函数的返回值去接这个错误信息
	_,err3 := chufa(2,0)		//这里可以先做判空处理，err3 != nil再使用c
	fmt.Println(err3)

	//panic显示调用
	fmta()
	fmtb(6)
	fmtc()
}

func chufa (a,b int) (c int,err error) {
	if(b == 0){
		err = errors.New("被除数不能为0")
	}else{
		c = a/b
	}
	return
}

func fmta(){
	fmt.Println("aaaa")
}

func fmtb(i int){
	//设置recover,这个时候就只有发生panic以后本函数里的不执行，不会导致程序奔溃，缺点是无法输出错误信息，需要手动判断是否为空，然后控制输出
	defer func() {
		//recover()
		//fmt.Println(recover())			//这里打印错误信息，但是没有错误的时候输出nil，这里应该判断一下
		if err := recover(); err != nil{	//这里没有错误的时候就不打印了
			fmt.Println(err)
		}
	}()
	//fmt.Println("bbbb")
	//panic("this is panic")			//panic显示调用，结果和老师讲的不太像
	var a [4]int
	a[i] = 111				//我的不会输出错误，但是会直接结束程序
	fmt.Println("bbbb")
}

func fmtc(){
	fmt.Println("cccc")
}