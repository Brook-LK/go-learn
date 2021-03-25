package hello //同一个目录报名必须一样，它放在main包中，项目文件必须放在src文件中

import (
	"fmt"		//控制面板输出
	"os"		//控制面板指令
)
//点操作
//import . "fmt"	//这种方法导入的包可以直接调用，如Println("点操作输出"),不常用，容易乱，分不清

//给包名起别名
import io "fmt" //给fmt起别名为io，调用io.Println()

//忽略包
//import _ "fmt"	//这里不调用fmt也可以编译

//查看环境路径
//cmd:go env

//同一个目录调用别的文件的函数无需应用，直接调用,如调用hellofunc.go中的函数
//不同目录调用，需要引入调用函数所在包，调用发放"包名.方法名"，其中调用方法名首字母必须大写，定义的时候首字母也必须大写

var global int

//变量定义及使用
func Main1() {
	//go代码行结尾可以不写分号;
	//定义常量,在go中定义或导入的包没有引用会报错
	//go 定义变量的时候可以不告诉定义的变量类型，它会自己判单属于什么类型
	fmt.Println("定义变量")
	var a1 int	//a1=0,通过var定义的变量必须有引用
	var a2 = 20;
	const a3 int = 3	//const也可以定义常量
	const a4 = 4
	var ( 			//var也可以批量定义变量
		a5 int
		b4 float32
	)
	fmt.Println(a1)
	fmt.Println("a2 = ",a2)	//go输出拼接方法
	fmt.Println(a3)
	//fmt.Println(a4)
	fmt.Println(a5)
	fmt.Println(b4)
	var b1 float32
	var b2 float64
	const b3 float32 = 1
	fmt.Println(b1)
	fmt.Println("b2 = ",b2)
	fmt.Println(b3)

	//自动推导类型
	fmt.Println("定义推导类型")
	c := 10
	fmt.Println(c)
	fmt.Printf("c = %T",c) //%T为返回类型+值int0

	//交换两个数
	a1,c = c,a1
	fmt.Println(c)

	//_为匿名变量，丢弃数据，没法后继调用。配合函数返回值使用才有优势，不用的直接用匿名变量
	a2,_ = a1,a2
	fmt.Println(a1)
	fmt.Println(a2)

	//多参数返回
	_,a2,a1 = test()
	fmt.Printf("b = %d, c = %d",a2,a1)

	//iota常量自动生成器,每行自动累加1，
	//iota给常量赋值使用，遇到const会重置为0
	const(
		d1 = iota   //0
		d2 = iota   //1
		d3 = iota   //2
	)
	const d = iota   //0
	fmt.Println(" ")
	fmt.Println("d1 = ",d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Printf("d = %d",d)

	//定义byte（字符）
	var ch byte
	ch = 97
	fmt.Printf(" %c,%d",ch,ch)
	ch = 'b'
	//大小写转换
	ch = 'b' - 32
	ch = 'B' + 32

	//字符串
	//字符串都是隐藏了一个'\0'
	var str string
	str = "hello"
	fmt.Printf("%c",str[0])  //h,字符串是由字符组成
	str1 := "world"
	var str3 string
	fmt.Println(str1)
	fmt.Println(len(str1))
	fmt.Println(str3)

	//复数类型
	var t complex64 //申明
	t = 2.1 + 3.14i
	fmt.Println(t)
	t1 := 2.1 + 3.14i //自动推导类型
	fmt.Printf("t type is %T\n",t1)		//转义换行
	//通过内建函数，取实部和虚部
	fmt.Println("real(t1) = ",real(t1),"imag(t1) = ",imag(t1))

	//常用%
	//%T输出格式，%c输出字符，%b输出二进制，%d输出10进制，%v使用默认格式输出，%s输出字符串，%f输出浮点型，%t输出bool类型

	//从键盘读数据
	var s6 string
	fmt.Println("请输入一串字符串")
	//fmt.Scanf("%s",&s6)	//阻塞等待用户输入,&号不能少
	fmt.Println("您输入的字符串为 ",s6)
	//fmt.Scan(&s6)
	fmt.Println("您输入的字符串为 ",s6)

	//类型转换，转换两个可以兼容的类型
	var flag bool		//在go中int不可以转bool时，0不是假，非0也不是真
	flag = true
	fmt.Printf("flag = %t\n",flag)
	var ch1 byte
	ch1 = 'a'
	var a7 int
	a7 = int(ch1)
	fmt.Println(a7)

	//类型别名
	type bigint int64	//给int64起个别名叫bigint
	type(
		mystring string
		myflout float32
	)
	var a8 bigint
	a8 = 10
	fmt.Printf("a8 type is %T\n",a8)  	//a8 type is main.bigint

	//运算符 + - * / % ++ --
	//关系运算符 == != < > <= >=
	if(a1 < a7){			//可以不带(),if a1 < a7
		fmt.Println( "a1<a7")
	}else if(a1 == a7){
		fmt.Println( "a1==a7")
	}else{
		fmt.Println( "a1>a7")
	}
	//逻辑运算符 ! && ||
	//位运算符  & | ^(异或) << >>
	//赋值运算符 = += -= *= /= %= <<= >>= &= ^= |=
	//其他运算符 &(取地址运算符) *(取值运算符)    例子：fmt.Scan(&a)
	//运算符优先级分7级,一元最高，二元运算方向均是从左至右
	//(^ !),(* / % << >> & &^),(+ - | ^),(== != < <= >= >),(<-),(&&),(||)

	//控制流程
	//if if-else if-(else if)-else...
	//if支持一个初始化语句
	if a9 := 9; a9 == 10{
		fmt.Println("if支持一个初始化语句,如在判断前初始化a9")
	}else if a9 := 10; a9 == 9{
		fmt.Println("hehe,这是什么逻辑")
	}else{
		fmt.Println("a9明明等于10")
	}
	//switch,像if一样支持一个初始化语句
	switch a10 := 12; a10 {
	case 1:
		fmt.Println("输入的数位1")
		//break		//默认带了break，不写也可以
		fallthrough	//不论当前判断是否成功都继续执行，不跳出,而且该判断后的东西无条件输出，不包括default中内容，如果成立，等于13的也会输出
	case 13,12,14:	//case后买你也可以带多个
		fmt.Println("输入的数位13")
	default :
		fmt.Println("输入的数未记录在案")
	}
	a11 := 13
	switch  {    //可以没有条件
	case a11 > 11:	//不像Java case后面也可以判断
		fmt.Println("a11 > 11")
	case a11 < 8:
		fmt.Println("a11 < 8")
	}
	//循环，go中没有while，do while循环，只有for循环和range迭代器
	a12 := 0
	for i:=1; i<=100; i++ {	//for后面不能加(),
		a12 += i
	}
	fmt.Println("a12 = ",a12)
	str2 := "abcdef"
	for i:=0;i<len(str2);i++{
		if i == 2{
			continue
		}else if i>4{
			break
		}
		fmt.Printf("str2[%d] = %c\n",i,str2[i])
	}
	goto FLAG
	for i,data := range str2 {
		fmt.Printf("str2[%d] = %c\n",i+10,data)
	}
	for i := range str2 {		//丢弃第二个参数，只返回下标i
		fmt.Printf("str2[%d] = %c\n",i+20,str2[i])
	}
	FLAG:
	for i,_ := range str2 {		//丢弃第二个参数，只返回下标i
		fmt.Printf("str2[%d] = %c\n",i+30,str2[i])
	}

	//跳转语句，break可用于for,switch,select ,而continue只能用于for循环
	//goto可以用于任何地方，但是不能垮函数使用
	goto FLAG1

	FLAG1:
		fmt.Println("这里为goto输出")

	//定义函数
	fmtPrint()
	fmtPrint1(a12)
	fmtPrint2(a12,a1)
	fmtPrint3(a12,str1)
	fmtPrint4(3,2,1,4,5)
	fmtPrint5(3,2,1,4,5)
	fmtPrint6(3,2,1,4,5)

	var a20 = fmtReturn()
	var a22 int
	a21 := fmtReturn()
	fmt.Println(a20)
	fmt.Println(a21+1)
	//fmt.Scan(&a20)
	a22,a20 = fmtReturn1(2)
	_,a20 = fmtReturn1(5)
	fmt.Printf("a22 = %d ,a20 = %d \n",a22,a20)
	a23,a24,_,a25 := fmtReturn2()
	fmt.Printf("a23 = %d ,a24 = %d ,a25 = %d \n",a23,a24,a25)
	a26 := fmtReturn3(4)
	fmt.Println(a26)
	fmt.Println(FmtReturn4(1,100))
	//给函数起别名，起别名前需要先定义函数类型其中包括形参和实参的类型
	var add FuncType1
	add = FmtReturn4
	fmt.Println(add(1,100))
	fmt.Println(fmtReturn5(2,100,add))

	//匿名函数，没有函数名的,函数定义，还需要调用，可以获取所在方法种的变量
	f1 := func(){	//自动推导类型
		fmt.Println(a26)
	}
	f1()
	type FuncType2 func()   	//定义函数类型
	var f2 = f1			//函数赋值+自动识别
	f2()
	//定义匿名函数并直接调用
	func(){
		fmt.Println(a26+1)
	}()		//后面的括号表示调用，但是这个方法只能在这里调用一次，因为别的地方取不到它
	func(a int){
		fmt.Println(a+2)
	}(a26)
	a27,a28 := func(i,j int)(max,min int) {		//有返回值的接法
		if i > j {
			max = i
			min = j
		}else{
			max = j
			min = i
		}
		return 
	}(2,5)
	fmt.Println(a27,a28)
	//闭包函数对外部函数的影响
	func() {
		a27++
		fmt.Println(a27)
	}()
	fmt.Println(a27)	//闭包函数调用外部变量会影响外部变量属性

	//函数在没有被释放前初始化数据是会保留的,只要闭包还在使用，它不关系是否超出作用域，只要还有闭包还在使用这些变量就还回会存在
	fmt.Println(fmtReturn6())	//1
	fmt.Println(fmtReturn6())	//1
	fmt.Println(fmtReturn6())	//1
	f3 := fmtReturn7()
	fmt.Println(f3())	//1
	fmt.Println(f3())	//4
	fmt.Println(f3())	//9
	fmt.Printf("fmtReturn7() is a %T\n",fmtReturn7())		//这样只是返回一个返回int值的函数
	fmt.Println(fmtReturn7()())		//再次执行返回的函数才可以，

	//defer的使用,延迟调用，在main将要关闭的时候调用，在没有遇到未定义defer的错误函数之前都可以调用到，如果是可能错误函数也定义了defer则都能输出
	defer fmt.Println("这是main函数通过defer调用的1")
	//defer 的执行顺序，后进的先出
	defer fmtReturn8(100,0)
	defer fmt.Println("这是main函数通过defer调用的2")
	//匿名函数和defer一起使用时，如果匿名函数有传参则传入当前参数值，如果直接调用则调用到的值为触发defer时的值
	defer func() {
		fmt.Println("1",a27)
	}()
	defer func(a int) {
		fmt.Println("2",a)
	}(a27)			//程序跑到这里的时候已经将a27的值传进了匿名函数
	a27 = 13
	fmt.Println("3",a27)

	//获取命令行参数
	//接受用户传递的参数，都是以字符串方式传递
	list := os.Args
	fmt.Println(len(list))
	for i,data := range list {
		fmt.Printf("指令第 %d 个参数为 %s\n",i,data)
	}

	//局部变量和全局变量，
	//局部变量，定义在大括号里的就是局部变量作用域为这个大括号内
	//全局变量，定义在函数外的就是全局变量,包括可以在main外面,但是只能在函数或者main函数中初始化赋值？
	global = 15
	fmt.Println("全局变量",global)
	//不同作用域同名变量，不同作用域允许定义同名变量，使用变量使用就近原则
	var global = 12
	fmt.Println("局部变量",global)

	io.Println("包别名输出")






}



