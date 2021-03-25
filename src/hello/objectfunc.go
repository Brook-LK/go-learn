package hello

import(
	"fmt"
)

func Print(){
	fmt.Println("this is hello.objectfunc")
}

//结构体方法，需要绑定某一种类型，绑定类型为接收者,面向对象一般调用该种函数。接收这就是传递的一个参数。不支持重载，只要接收者类型不同就可以，但是不能相同接收者不同类型的入参
func (i myint) Add(other myint) myint{
	return i+other
}

func (stu *Student1) PerToStu(per Persion) {		//其中接收者类型不能直接是指针类型，可以是*Student1这样写的,因为指针中你不知道存了什么
	stu.Persion = per
}

//下面这个方法会报错因为方法名与上一个相同
//func (stu *Student1) PerToStu(per int) {
//	stu.Persion.Id = per
//}

func (stu Student1) PerToStu1(per Persion) {
	stu.Persion = per
}

//方法继承
func (per *Persion) SetPersion(id int,name string){
	per.Id = id
	per.Name = name
	fmt.Println("Persion.SetPersion")
}

//方法重写
func (stu *Student1) SetPersion(id int,name string){
	stu.Id = id
	stu.Name = name
	fmt.Println("Student1.SetPersion")
}


