package hello

import (
	"fmt"
)

func Lucky()  {
	fmt.Println("请在接下来的提示中输入'喜欢'或者'不喜欢'!  ^-^ ")
	fmt.Println("你比较喜欢吃米饭?")
	var input string
	fmt.Scan(&input)
	fmt.Printf("我知道了，你%s吃米饭\n",input)
	for {
		fmt.Println("你喜欢LK?")
		var love string
		fmt.Scan(&love)
		if love == "喜欢" {
			fmt.Printf("我知道了,你%sLK,LK也%s你\n",love,love)
			goto LOVE
		}else{
			fmt.Println("哼，我没有收到我想看到的消息，请重新输入！")
		}
	}

	LOVE:
		fmt.Println("一个人要保护好自己哦，出门记得戴口罩，回家记得消毒！")
	var by string
	fmt.Scan(&by)
}
