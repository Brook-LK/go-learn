package hello

import (
	"fmt"
	"time"
	"math/rand"
)
func GuessNum(){
	num := getGuess()
	fmt.Println(num)
	fmt.Println("start hello.GuessNum,place input a number between 0~100")
	var input int
	fmt.Scan(&input)
	var temp int
	if input > num{
		temp = 0
	}else{
		temp = 100
	}
	for i := 1;;i++{
		fmt.Scan(&input)		//这里需要多家判断以免输入不规范的数据
		if num == input{
			fmt.Printf("你真厉害第 %d 就猜对了",i)
			break
		}else if num > input{
			fmt.Printf("%d < randNum < %d\n",input,temp)
			temp = input
		}else{
			fmt.Printf("%d < randNum < %d\n",temp,input)
			temp = input
		}
	}

}

func getGuess()(a int){
	rand.Seed(time.Now().UnixNano())
	a = rand.Intn(100)
	return
}
