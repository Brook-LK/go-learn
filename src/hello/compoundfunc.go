package hello
//复合类型函数
import (
	"fmt"
)

func swap1(a,b int)  {
	a,b = b,a
	fmt.Printf("a = %d ,b = %d\n" ,a,b)
}

func swap2(a,b *int)  {
	a,b = b,a
	fmt.Printf("a = %d ,b = %d\n" ,*a,*b)
}

func swap3(a,b *int)  {
	*a,*b = *b,*a
	fmt.Printf("a = %d ,b = %d\n" ,*a,*b)
}

func maopao(args [6]int) (args1 [6]int) {
	n := len(args)
	for i := 0;i<n-1;i++{
		for j := 0;j<n-1-i;j++{
			if args[j] > args[j+1]{
				args[j],args[j+1] = args[j+1],args[j]
			}
		}
	}
	args1 = args			//数组做函数参数，它是值传递，不是对象传递
	return
}

func maopao1(args *[6]int) {	//采用地址传递
	n := len(args)			//这里也可以是len(*args)这么写
	for i := 0;i<n-1;i++{
		for j := 0;j<n-1-i;j++{
			if args[j] > args[j+1]{			//这里args[j]取到的对象就是改数组地址下第j个数，不需要*arg[j]
				args[j],args[j+1] = args[j+1],args[j]
			}
		}
	}
}

func maopao2(args []int) {	//切片采用引用传递
	n := len(args)
	for i := 0;i<n-1;i++{
		for j := 0;j<n-1-i;j++{
			if args[j] > args[j+1]{
				args[j],args[j+1] = args[j+1],args[j]
			}
		}
	}
}

func deleteMapKey(m map[int]string){
	delete(m,111)
}

func changeStudent (s Student) {	//结构体作为函数参数为值传递，无法改变实参
	s.name = "change1"
	fmt.Println(s)
}

func changeStudent2 (s *Student) {	//结构体地址作为函数参数为引用传递
	s.name = "change2"
	fmt.Println(*s)
}