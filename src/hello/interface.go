package hello

import(
	"fmt"
)

func Interface(){
	fmt.Println("start Interface")
	//定义接口类型的变量
	var man Man
	//只要实现了此接口方法的类型，那么这个类型（接收者类型）的变量就可以给man赋值(感觉就是把内容封装了一下，类似于java中view&model，比如这里可以是输出到前台)
	var str string = "124"
	stu1 := Student1{Persion{1,"kk","男"},"12","name",3,2,&str}
	man = &stu1
	man.PrintStudent1()
	//man.PrintStudent2()		//接口转换，超集可以转换调用子集的函数，但是子集不能调用超级的函数

	//多态
	DPrintStudent(&stu1)
	//创建一个切片，里面放的是接口
	s := make([]Man,1)		//s := make([]interface{},n)万能接口，这样可以存储不知Man类型的接口
	s[0] = &stu1
	for i := 0;i< len(s);i++{
		DPrintStudent(s[i])		//两种方法调用接口
		s[i].PrintStudent1()
	}

	//接口继承，StudentI继承了Man
	var stui StudentI
	stui = &stu1
	stui.PrintStudent1()		//StudentI继承过来的方法
	stui.PrintStudent2()		//StudentI自己独有的方法

	//接口转换，超集可以转换为子集，反过来不可以，就像Student中没有实现PrintStudent2，但是可以调用PrintStudent2(),反过来不可以
	s[0] = stui		//超集转换为子集
	//stui = s[0],会报错，子集不能转换为超集

	//空接口万能类型，保存任意类型的值,所有类型都实现了这样一个方法，比如Print方法就是：func Println(a ...interface{})
	var W interface{} = 1
	fmt.Println("W = ",W)
	W =  "hello"
	fmt.Println("W = ",W)

	//类型查询，断言
	if value,ok := W.(string);ok == true { //value返回的是值，ok返回判断真假,这里的W好像不能是已经明确的基本类型，一般是接口中的
		fmt.Println("W = ",W,";value = ",value)
	}
	switch value := W.(type) {			//和通过if断言写法有点不一样
	case int:
		fmt.Println("W = int =",value)
	case string:
		fmt.Println("W = string =",value)
	}

}

