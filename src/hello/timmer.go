package hello
//定时器
import (
	"fmt"
	"time"
)

func Timmer()  {
	fmt.Println("this is helle.Timmer")
	//创建一个定时器，设置时间为2s，2s后往time通道写数据
	timer := time.NewTimer(5*time.Second)		//只会执行一次
	timer.Reset(1*time.Second)		//重新设置timer
	fmt.Println("当前时间为",time.Now())
	//timer.Stop()		//停止定时器，
	t := <-timer.C
	fmt.Println("t = ",t)
}

func Ticker(){
	ticker := time.NewTicker(2*time.Second)			//Ticker为循环定时往chan写数据,方法和timer一样
	fmt.Println("当前时间为",time.Now())
	for{
		t := <-ticker.C
		fmt.Println("t = ",t)
	}
}

//select主要用来监听channel管道上的数据流动，和switch很像，其中case必须为一个io操作
func getChannel(ch chan int){
	select {
	case n := <-ch:
		//如果ch成功读取到数据则执行该case代码块
		fmt.Println(n)
	case ch <- 2:		//如果上面和下面的都能执行，则选择其中一条执行 
		//如果ch成功写入则执行此case代码块
		fmt.Println("成功写入ch")
	default:
		fmt.Println("select default执行")
	}
}
