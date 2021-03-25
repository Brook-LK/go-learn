package hello

import(
	"fmt"
)

//定义接口类型
type Man interface {			//子集
	//该方法由别的类型（自定义类型）实现，只有申明，没有实现
	PrintStudent1()
}

//多态,定义一个普通，函数的参数为接口类型（相当于包装）
func DPrintStudent(man Man){
	man.PrintStudent1()
}

//用于实现Man接口中的PrintStudent1方法
func (stu *Student1) PrintStudent1 () {
	fmt.Println("this is PrintStudent1 print student",stu)
}

func (per *Persion) PrintStudent1 () {
	fmt.Println("this is PrintStudent1 print persion",per)
}

//接口的继承
type StudentI interface {			//超集
	Man				//匿名字段继承了Man
	PrintStudent2()			//也可以拥有自己单独的接口方法，和结构体继承类似
}

func (per *Persion) PrintStudent2 () {
	fmt.Println("this is PrintStudent2 print persion",per)
}

//func (stu *Student1) PrintStudent2 () {
//	fmt.Println("this is PrintStudent2 print student",stu)
//}




