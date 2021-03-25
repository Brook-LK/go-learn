package hello

import(
	"fmt"
)

//封装：通过方法实现，不支持重载
//继承：通过匿名字段实现
//多态：通过接口实现

//专用名词，值语义（值传递）和引用语义（引用传递）

func Object(){
	fmt.Println("there is hello.Object")
	var str string = "124"
	stu1 := Student1{Persion{1,"kk","男"},"12","name",3,2,&str}
	var stu2 Student1 = Student1{Persion{1,"kk","男"},"13","name2",4,6,&str}
	fmt.Println(stu1)
	fmt.Println(stu2)
	fmt.Printf("stu2 is %v\n" ,stu2)
	fmt.Printf("stu2 is %+v\n" ,stu2)		//%+v可以显示更详细的信息
	//指定成员初始化，没有指定的为默认值
	stu3 := new(Student1)
	*stu3 = Student1{Persion:Persion{Id:3,Name:"ll"}}
	stu3.Name = "666"		//同名属性，这里只有外面的被修改了
	stu3.Persion = Persion{2,"lll","www"}			//整体匿名对象赋值
	fmt.Println(*stu3)
	fmt.Println(stu3.Name,stu3.myint)
	stu3.Persion.Name = "777"
	stu3.myint = 15			//基础匿名字段赋值
	fmt.Println(*stu3)

	//简单方法调用
	//var myint1 myint		//myint1 := 2会被认为是int型
	//myint1 = 2
	var myint2 myint = 3
	fmt.Println(myint2.Add(2))

	//结构体调用方法
	var per1 Persion = Persion{1,"hello","temp"}
	var per2 Persion = Persion{1,"hello2","temp2"}
	var stu4 Student1
	stu4.PerToStu(per1)
	fmt.Printf("stu4 = %+v\n",stu4)
	(&stu4).PerToStu(per2)			//这两种方法都可以，就是如果是结构体类型的是可以不用&来区分的（好像引用传递的都可以）
	fmt.Printf("stu4 = %+v\n",stu4)

	//结构体变量是一个指针变量，它能够调用哪些方法，这些方法就是一个集合，简称方法集（会自动先把指针或内存转换为对应的接收者，然后再调用方法，所以调用的时候不需要考虑）
	(&stu4).PerToStu(per1)
	(&stu4).PerToStu1(per1)		//既可以调用接收者为指针的又可以调用接收者为结构体指针内存的,没有什么区别，内部会先把他转换成内存或指针

	//普通变量的方法集，就像上面给的赋值一样，先赋值个变量，后传指针&str

	//方法的继承，通过匿名字段的方式继承了匿名字段所拥有的方法
	stu4.Persion.SetPersion(12,"from P")		//被重写后希望调用到原来的方法，得这么调用，称为显式调用继承方法
	fmt.Printf("stu4(SetPersion) = %+v\n",stu4)

	//方法重写
	stu4.SetPersion(13,"from S P")			//重写后调用到的就只是这个结构体的这个方法，类似于就近原则，而且如果有同名属性，注意被操作的属性也是就近原则
	fmt.Printf("stu4(SetPersion) = %+v\n",stu4)

	//方法值
	sFunc := stu4.SetPersion
	sFunc(14,"from S P1")		//有点类似之前的占位		//等价于stu4.SetPersion(14,"from S P1")
	fmt.Printf("stu4(SetPersion) = %+v\n",stu4)

	//方法表达式
	sFunc1 := (*Student1).SetPersion
	sFunc1(&stu4,15,"from S P2")		//通过方法表达式调用的时候第一个传值为接收者
	fmt.Printf("stu4(SetPersion) = %+v\n",stu4)
}

type Persion struct {
	Id int
	Name string
	Sex string
}

type Student1 struct {
	Persion			//匿名字段，没有名字，只有类型，相当于把Persion的属性都放到Student中，实现代码复用，也即继承
	ClassName string
	Name string 	//同名字段，各自赋值，相当于重写了，就近原则（如果能在本作用域找到此成员变量就只作用于本作用域）
	myint 			//基础类型的别名匿名字段
	int				//基础类型的匿名字段
	*string			//结构体指针类型,指针内存可以是任何类型,如果是基础类型得先创建类型对象
}


type myint int
