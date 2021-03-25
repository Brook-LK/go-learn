package hello
//goroutine,go的协程，比线程更少，go通过通信来共享内存，而不是共享内存来通信，通过channel实现同步(无缓冲的channel，在没有读(获取到数据),或写不进去数据的时候都会被阻塞)
import (
	"fmt"
	"runtime"
	"time"
)

//var ch = make(chan int)    //通过channel实现同步，这里管道中的类型可以是其他类型，无缓冲的channel，相当于make(chan int,0)
var ch = make(chan int,100)			//创建有缓存的通道，可以通过len(ch),cap(ch)查看容量和数据，当管道满了以后也会阻塞
//管道分为双向管道和单向管道，上面的为双向管道var ch1 chan<- int = ch,(只写管道)，var ch1 <-chan int = ch,(只读管道),数据存放在ch中

func Goroutine(){			//主协程，如果主协程执行太快可能导致子协程没有被调用
	n := runtime.GOMAXPROCS(8)	//n为电脑逻辑处理器，1为指定为单核运行，最大可谓n
	fmt.Println(n)
	go fmtPrintln()				//新建一个协程（子协程），通过go新建协程（线程）
	//for {
	for i := 0;i<100;i++ {		//发现主协程退出后子协程也退出，可以写个死循环，当子协程执行完毕后通过goto去控制
		runtime.Gosched()		//主协程让出时间片，让其他协程先执行，不然每次都是先输出"this is hello.Goroutine",加了i变量感觉好像不起作业
		time.Sleep(time.Second)			//延时一秒
		fmt.Println("this is hello.Goroutine")		//默认先执行主协程，所以每次都是这句话先输出
		if i >= 2 && i < 10 {
		//if i == 2 {
			fmt.Printf("i = %d 管道中有数据开始执行fmtPrintln()\n",i)
			ch <- i			//给管道ch写数据
			if i>=10 && len(ch) == 0 {		//关闭的时候得正常关闭
				close(ch)		//关闭channel,通道有数据的时候不能关闭,关闭后是没法再像管道发送数据
			}
		}
	}

	//核数测试（并发测试）
	/**
	for {
		go fmt.Print(1)			//runtime.GOMAXPROCS(8)中数越大，0，1输出间断越小
		fmt.Print(0)
	}*/


}

func fmtPrintln() {
	//n := <-ch 		//从管道取数据，接收，如果没有数据就会阻塞，这个时候就不会执行后续内容，知道ch有数据，ch也可以放到主进程中
	//fmt.Printf("n = %d\n",n)		//拿到管道中的数据，实现数据交换
	for {
		//n := <-ch			//无缓冲的channel，数据取出来后就没有了，后面的代码只有ch中有值的时候才能被执行,ch没有被读之前写ch的地方也会被阻塞
		//fmt.Printf("n = %d\n",n)
		if n , ok := <- ch ; ok == true {
			time.Sleep(3*time.Second)
			fmt.Printf("n = %d\n",n)
			fmt.Println("this is hello.fmtPrintln")
		}else{
			fmt.Println("通道关闭")
			return
		}
		//time.Sleep(3*time.Second)
		//fmt.Println("this is hello.fmtPrintln")
		//runtime.Goexit()	//终止所在协程
	}
}

func truech (ch chan int) (bo bool){
	if n , ok := <- ch ; ok == true {
		fmt.Printf("n = %d\n",n)
	}else{
		fmt.Println("通道关闭")
	}
	return
}

//遍历管道中的数据
func rangeChan (ch <-chan int) {		//只读，chan可以隐式的转换为只读或只写chan，chan为引用传递
	for num := range ch{
		fmt.Println(num)	//遍历管道，后管道中还有数据么，和<-ch有没有区别
	}
}