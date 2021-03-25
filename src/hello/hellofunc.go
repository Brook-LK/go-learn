package hello

import "fmt"

//定义函数,函数也是一种数据类型，也可以通过type改名
//写法
//func funcName(/*"入参"*/)(o1 int,o2 string/*返回参数及类型*/){
//	/*函数体*/
//}
func test()(a,b,c int){
	return 1,2,3
}

//无参函数,函数名后的第一个括号为入参，第二个括号为输出，如果没有第二个括号则表示没有输出
func fmtPrint(){
	fmt.Println("这是无参函数的输出")
}

//有参无返回值函数
func fmtPrint1(a int){
	fmt.Println(a)
}

func fmtPrint2(a,b int){	//如果入参，a,b是同一种类型，这样写
	fmt.Println(a)
	fmt.Println(b)
}

func fmtPrint3(a int,b string){	//如果入参，a,b是同一种类型，这样写
	fmt.Println(a)
	fmt.Println(b)
}

//不定参数类型, ...int 意思为，类似int的类型;   ...type ,不定参数类型,不定参只能放在形参的最后一个，调用函数的地方传入的参数叫实参
func fmtPrint4(args ...int){	//传递的参数可以是0或多个,
	fmt.Println("args size is ",len(args))
	for i := 0;i < len(args); i++{
		fmt.Printf("args[%d] = %d\n",i,args[i])
	}
}

func fmtPrint5(a int,args ...int){	//传递的参数可以是0或多个,
	fmt.Println("定参为 ",a)
	fmt.Println("不定参个数为",len(args))
}

func fmtPrint6(args ...int){
	fmtPrint7(args...)		//调用其他函数args...表示把fmtPrint6的形参args全部传给fmtPrint7
	fmtPrint7(args[1:4]...)	//从下标为1的开始到下标为4但不包含4的参数
	fmtPrint7(args[3:]...)	//从下标为3的开始到结尾的参数
}

func fmtPrint7(args ...int){
	fmt.Println("args7 size is ",len(args))
	for _,data := range args{
		fmt.Print(data,",")
	}
	fmt.Println(";")
}

//无参有返回值调用
func fmtReturn() int {
	var a = 3
	return a	//带参数返回
}

func fmtReturn1(a int) (b,c int){
	b = a+1
	c = a+12
	return		//前面定义了的后面可以直接返回
}

func fmtReturn2() (a,b,c,d int){
	a,b,c,d = 3,4,6,7	//多赋值
	return
}

//函数名大写表示是共有的，小写表示是私有的，有返回值的函数必须有return

func fmtReturn3(a int) (b int){		//递归调用
	fmt.Println(a)
	if a > 0 {
		a--
		b = fmtReturn3(a)
	}else {
		b = 12
	}
	return
}

//递归累加1,a为起始数，b为终止数，每次a+1累加到c
func FmtReturn4(a,b int) (c int){
	if a <= b{
		c = FmtReturn4(a+1,b) + a
	}
	return
}

//定义函数类型，用于给函数起别名
type FuncType1 func (int,int) (int)

//回调函数，函数有一个参数是函数类型，这个函数就是回调函数,类似于Java种的泛型只不过这里是方法
func fmtReturn5(a,b int ,myfunc FuncType1) (c int){
	c = myfunc(a,b)
	return
}

func fmtReturn6() int {
	var a29 int
	a29++
	return a29 * a29
}

func fmtReturn7() func() int {		//返回一个函数，这个函数返回一个int值。
	var a29 int
	return func() int{				//返回了一个int
		a29++
		return a29 * a29
	}
}

func fmtReturn8(a,b int){
	fmt.Println(a/b)
}

func FmtReturn9(a,b int) (c int){
	if a <= b{
		c = FmtReturn9(a+1,b) + a
	}
	return
}

func init() {
	fmt.Println("this is hello.hellofunc.init")
}