package hello
//复合类型，pointer,指针，array，数组，slice切片，map字典，struct结构体
import (
	"fmt"
	"math/rand" //随机数
	"time"      //获取本地时间
)

func Compound(){
	//每个变量有两层含义，变量的内存，变量的地址默认值nil
	var a int = 10
	fmt.Printf("a =  %d\n",a)	//变量的内存（数据）
	fmt.Printf("a =  %v\n",&a)	//变量的地址，指针
	//保存某个变量的地址，需要指针类型，     *int 保存int的地址，**int保存*int的内存（数据）
	//申明（定义), 定义只是特殊的申明
	//定义个变量p，类型位*int
	var p *int
	p = &a
	fmt.Printf("p =  %v\n",p)
	fmt.Printf("&a =  %v\n",&a)
	*p = 666 //*p操作的不是p的内存，是p所指向的内存（就是a），也就是说*p = 等价于 a = ，*p操作的其实是p指向地址内存的数据
	fmt.Printf("*p =  %v\n",*p)
	fmt.Printf("a =  %v\n",a)
	fmt.Printf("&a =  %v\n",&a)
	fmt.Printf("a =  %d\n",a)
	//不要操作没有初始化的地址，会报错
	var q *int
	//*q = 100		//这里没有初始化直接赋值会报错
	//需要初始化后才能用
	q = new(int)
	*q = 100
	fmt.Printf("*q = %d\n",*q)		//go语言只管初始化，语言会自己做内存处理
	m := new(int)
	*m = 101
	fmt.Printf("*m = %d\n",*m)

	//普通变量做函数参数，go复制传递，不是指针传递（引用传递）
	a1,a2 := 10,20
	swap1(a1,a2)
	fmt.Println(a1,a2)
	swap2(&a1,&a2)
	fmt.Println(a1,a2)
	swap3(&a1,&a2)			//执行前后地址不变，同时也能看出a,b = b,a等价于 &a,&b = &b,&a普通变量做函数参数的时候作用域外地址不变
	fmt.Println(a1,a2)

	//数组
	var ids [4]int		//定义数组个数必须是常量,[4]int 和 [10]int 是不同的类型
	fmt.Println(ids)
	for i,_ := range ids {		//这里不能用i,data
		ids[i] = i
	}
	fmt.Println(ids)
	fmt.Printf("ids is %T\n",ids)
	var ids2 [5]int = [5]int{1,2,3,4,5}		//定义同时初始化
	fmt.Println(ids2)
	ids3 := [6]int{1,2,3,4}		//部分初始化
	fmt.Println(ids3)
	ids4 := [6]int{1:2,3:4}		//指定下标初始化
	fmt.Println(ids4)
	//多维数组
	var ids5 [2][3]int
	fmt.Println(ids5)
	for i := 0;i<2;i++{
		for j := 0;j<3;j++{
			ids5[i][j] = i+j
		}
		fmt.Println(ids5[i])
	}
	fmt.Println(ids5)
	ids6 := [2][3]int{
		{5,6,7},{8,9,0},
	}
	fmt.Println(ids6)
	//数组比较，只支持同一类型的比较，只支持==和！=
	fmt.Println(ids3 == ids4)
	fmt.Println(ids3 == ids3)
	var ids7 [6]int
	ids7 = ids3
	ids3[4] = 5
	fmt.Println(ids7)

	//随机数，设置种子，产生随机数
	//rand.Seed(666)	//设置种子,如果种子参数一样，每次运行产生的随机数都一样
	rand.Seed(time.Now().UnixNano())		//以当前时间作为种子参数
	for i := 0;i<len(ids7);i++{
		//fmt.Println(rand.Int())		//rand.Intn(100) //限制随机数在100以内
		ids7[i] = rand.Intn(100)
	}
	fmt.Println(ids7)
	//for i := 0;i<len(ids7)-1;i++{
	//	for j := 0;j<len(ids7)-1-i;j++{
	//		if ids7[j] > ids7[j+1]{
	//			ids7[j],ids7[j+1] = ids7[j+1],ids7[j]
	//		}
	//	}
	//}
	ids8 := maopao(ids7)
	fmt.Println(ids8)
	fmt.Println(ids7)
	maopao1(&ids7)
	fmt.Println(ids7)

	//切片，弥补了数组的不可扩充，其实只是
	ids9 := []int{1,2,3,6,7,8,9}	//不指定长度的就是切片，也可以这样定义 [...]int{1,2,3,6,7,8,9}
	fmt.Println(ids9)
	fmt.Println(len(ids9))
	ids10 := ids9[1:4:6]		//表示切数组[1,4),其中容量为6-1个，长度为4-1个
	fmt.Println(ids10)
	fmt.Println(len(ids10))		//长度
	fmt.Println(cap(ids10))		//容量
	//创建切片
	ids11 := []int{}			//1、自动推导类型
	fmt.Printf("len = %d,cap = %d\n",len(ids11),cap(ids11))
	ids11 = append(ids11, 11)		//给切片末尾追加一个成员
	fmt.Printf("len = %d,cap = %d\n",len(ids11),cap(ids11))
	ids12 := make([]int,2,5)	//2、借助make函数，make(slice,len,cap)
	fmt.Printf("len = %d,cap = %d\n",len(ids12),cap(ids12))
	fmt.Println(ids12)
	ids13 := make([]int,2)	//3、借助make函数，make(slice,len),没有指定容量的时候容量和长度是一样的都是2
	fmt.Printf("len = %d,cap = %d\n",len(ids13),cap(ids13))
	//切片的截取，元素的操作和数组一样
	//ids14 := ids9[:]	//把ids9复制给ids14
	//ids15 := ids9[:4]	//等价于ids15 := ids9[0:4:4]
	//ids16 := ids9[4:]	//等价于ids16 := ids9[4:len(ids9):cap(ids9)]
	//切片和底层数组的关系
	ids10[1] = 666
	fmt.Println(ids9)		//修改切片后发现被切片的也变了，所以切片截取类似于引用一样
	fmt.Println(ids10)
	ids17 := ids10[1:4:5]
	fmt.Println(ids17)		//虽然ids17截取的是ids10，但是ids10是从ids9来的，就相当于ids17从ids10的第二位开始截取，而截取的内容是从ids9来的，也就是截取只是引用了
	ids17[2] = 777
	fmt.Println("ids9",ids9)
	//切片函数
	fmt.Printf("ids17 len = %d,cap = %d\n",len(ids17),cap(ids17))
	ids17 = append(ids17, 1)
	fmt.Printf("ids17 len = %d,cap = %d\n",len(ids17),cap(ids17))
	fmt.Println("ids9",ids9)		//给ids17追加后如果超过ids17的容量，ids17会进行扩容，且源不变，如果没有超出容量则覆盖被截取切片的下一位，长度与容量不变
	//所以修改截取切片的时候得谨慎
	fmt.Printf("ids9 len = %d,cap = %d\n",len(ids9),cap(ids9))
	ids17 = append(ids17, 2)		//扩容会后容量会变成原来的2倍
	ids17 = append(ids17, 3)
	ids17 = append(ids17, 4)
	ids17 = append(ids17, 5)
	fmt.Println("ids17",ids17)
	fmt.Printf("ids17 len = %d,cap = %d\n",len(ids17),cap(ids17))
	fmt.Println("ids9",ids9)
	fmt.Printf("ids9 len = %d,cap = %d\n",len(ids9),cap(ids9))
	ids17 = append(ids17, 6)
	fmt.Println("ids17",ids17)
	fmt.Printf("ids17 len = %d,cap = %d\n",len(ids17),cap(ids17))
	//copy 这个有点奇怪看执行结果吧
	copy(ids17,ids9)		//把ids17拷贝给ids9,被截取的没法拷贝
	fmt.Println("ids9",ids9)
	ids19 := []int{1,2,3,4,5}
	ids18 := []int{6,6,6,6,6,6,6,6,6,6,6,6}
	copy(ids19,ids18)
	fmt.Println("ids19",ids19)		//ids19 [6 6 6 6 6]
	//切片做函数参数的时候是引用传递
	maopao2(ids9)
	fmt.Println("ids9 ",ids9)

	//map,字典，映射
	//定义：map[keyType]valueType,keyType是唯一的，切片和函数是不能作为keyType的
	//map只有len，没有cap
	map1 := map[int]string{
		111:"make",
		222:"set",
	}
	fmt.Println(map1)
	fmt.Println(len(map1))
	map2 := make(map[int]string)
	map2[111] = "make"
	map2[111] = "go"		//有就修改value，没有就追加
	fmt.Println(map2)
	fmt.Println(map2[111])
	map3 := make(map[int]string,10)	//可以指定长度为10,但是这个容量没有什么做用,好处是提前分配好了容量
	fmt.Println(map3)
	fmt.Println(len(map3))
	//map循环遍历
	for key,value := range map1{
		fmt.Printf("map1[%d] = %s\n",key,value)
	}
	//判断一个key是否存在
	key := 111
	value,temp := map1[key]		//value为key对应的值，不存在的时候value为空、temp为false，存在时候value为值，temp为true
	if temp {
		fmt.Printf("key为%d存在于map1中，值为%s\n",key,value)
	}else{
		fmt.Printf("key为%d不存在于map1中\n",key)
	}
	//map删除key
	delete(map1,11)		//也可以删除map中原理不存在的，这个时候没有影响
	fmt.Println(map1)
	//map做函数参数，引用传递
	deleteMapKey(map1)
	fmt.Println(map1)

	//结构体，类似于Java中的实体
	//顺序初始化，每一个成员都必须初始化
	var stu1 Student = Student{1,"lk",2}
	fmt.Println(stu1)
	//指定成员初始化，没有初始化的为默认值
	stu2 := Student{id:1}
	fmt.Println(stu2)
	//结构体指针变量初始化
	var stu3 *Student = &Student{1,"lk",2}
	fmt.Println(*stu3)
	fmt.Printf("stu2 is a %T\n",stu2)
	//结构体成员使用，需要.运算符
	fmt.Println(stu1.name)
	stu1.name = "lkk"
	fmt.Println(stu1.name)
	//指针变量成员，指针有合法指向后才能操作
	var stu4 *Student
	stu4 = &stu1
	stu4.id = 2
	(*stu4).name = "lk"		//stu4.name和(*stu4).name完全等价
	fmt.Println(*stu4)
	//通过new来申请一个结构体
	stu5 := new(Student)		//这里是申请了一块内存，拿到的是地址
	stu5.name = "hello"
	fmt.Println(stu5)
	//比较，可以用==和!=比较，但是不能用>和<比较
	*stu5 = stu1			//同类型的两个结构体可以相互赋值
	if *stu5 == stu1{
		fmt.Println("*stu5 == stu1")
	}else{
		fmt.Println("*stu5 != stu1")
	}
	//作为函数参数为值传递，无法改变实参
	changeStudent(*stu5)
	fmt.Println(*stu5)
	changeStudent2(stu5)
	fmt.Println(*stu5)

}


//结构体，类似于Java中的实体，go中没有public,private,protected之分，只需要按照规定命名通过package.funcName或package.typeName调用（被调用的对象首字母必须大写，小写的不可见）就可以
type Student struct {
	id int		//这里没有var		//如果属性想在其他地方用也得改成大写，所以一般结构体首字母大写，属性首字母也大写
	//Id int
	name string
	age int
}