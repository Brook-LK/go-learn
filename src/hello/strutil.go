package hello

import (
	"bufio"         //缓存
	"encoding/json" //json转换
	"fmt"
	"io"
	"os"      //文件操作
	"regexp"  //正则表达式
	"strings" //字符串工具包
)

func StrUtil()  {
	fmt.Println("this is hello.StrUtil")
	//操作：strings.Contains,Join,Index,Repeat,Replace,Split,Trim,Fields     转换：strconv.Append,Format,Parse
	fmt.Println(strings.Fields("hello world 你好 大家好") )
	fmt.Println(strings.TrimSpace("     hello world 你好 大家好     ") )

	//[]byte转换为string   string([]byte)

	//正则表达式
	flag,err := regexp.MatchString(`a.b`,"abc adc abb acb")
	if err == nil {
		fmt.Println(flag)
	}else {
		fmt.Println(err)
	}
	//``包含的位原生字符串，遇到\也不会转义

	//Json
	//可以通过结构体和map生成json
	//结构体生成json的时候成员首字母必须大写
	tech1 := Teacher{"张三",18,28,38}
	tech2 := Teacher{"李四",19,29,39}
	teachs := []Teacher{tech1,tech2}
	//编码对象生成json文本
	json1,err2 := json.Marshal(teachs)
	//格式化编码 json.MarshalIndent(teachs,"","	")		//第二个位tab键
	if err2 == nil{
		fmt.Println(string(json1))
	}else {
		fmt.Println(err2)
	}
	//通过map生成json
	m := make(map[string]interface{},4)
	m["张"] = 2
	m["李"] = "三"
	m["zhou"] = Teacher{"李四",19,29,39}
	m["wu"] = teachs
	//json2,err3 := json.MarshalIndent(m,"","	")
	json2,err3 := json.Marshal(m)
	if err3 == nil{
		fmt.Println(string(json2))
	}else {
		fmt.Println(err3)
	}
	//解析json-struct
	var teac3 Teacher
	jsonStr := `{"Name":"张五","no":15,"Phone":"35"}`
	err4 := json.Unmarshal([]byte(jsonStr),&teac3)
	if err4 == nil {
		fmt.Println(teac3)
	}else {
		fmt.Println(err4)
	}
	//解析json-map
	m2 := make(map[string]interface{},4)
	err5 := json.Unmarshal(json2,&m2)
	if err5 == nil {
		fmt.Println(m2)			//反解析解析出来的都是map了，没有数组之分，成为map套map了，这里拿到的数据得类型断言，然后才能用数据，新版本的不用
	}else {
		fmt.Println(err5)
	}
	str1 := m2["李"]
	fmt.Printf("%T %s \n",str1,str1)

	//文件操作，文件分类：设备文件（屏幕，键盘），磁盘文件（文本文件，二进制文件）
	file,_ := os.Create("./hello.txt")		//默认当到当前调用目录下，如果文件不存在则新建，如果文件存在则清空文件并重新创建
	//file2,_ := os.Open("hello.txt")			//打开文件，但是是只读的方式。都调用了OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	file.WriteString("hello,writeString\n")
	//file2.Read()
	//file.Write("hello,writeString")			//WriteAt,带At的为在指定位置读或者写
	{
		var buf string
		for i := 0;i<10;i++{
			buf = fmt.Sprintf("i = %d\n",i)		//把i = i\n写入buf字符串中
			file.WriteString(buf)
		}
	}
	file.Close()				//关闭文件  或者defer file.Close()
	//os.Remove("./hello.txt")		//删除文件，没有关闭的文件是无法删除的
	//os.Stdout.Close()			//关闭输出
	os.Stdin.Close()			//关闭输入
	var a string
	fmt.Scan(&a)
	fmt.Println(a)
	os.Stdout.WriteString("are you ok\n")		//往屏幕输出东西
	{
		//复制文件
		file3,_ := os.Create("./word.png")
		//b,_ := ioutil.ReadFile("./hello.png")		//获取文件数据[]byte,文件结尾为io.EOF
		d,_ := os.Open("./hello.png")
		{
			//一边读，一边写
			c := make([]byte ,1024*2)		//2k
			for {
				n,err6 := d.Read(c)		//每次读c这么大
				if err6 != nil || err6 == io.EOF {
					fmt.Println("err6 = ",err6)
					break
				}
				file3.Write(c[:n])

			}

		}
		//file3.Write(b)
		file3.Close()
	}
	//行读取
	{
		file4,err7 := os.Open("./hello.txt")
		if err7 != nil {
			fmt.Println("err7 = ",err7)
		}
		r := bufio.NewReader(file4)		//新建一个缓冲区把内容先放在缓冲区
		for {
			buf,err8 := r.ReadBytes('\n')	//遇到\n结束读取,但是\n也读进去了
			if err8 != nil{
				if err8 == io.EOF {				//文件已经结束
					fmt.Println("文件已复制完成")
					break
				}
				fmt.Println("err8 = ",err8)
				break
			}
			fmt.Println("buf = ",string(buf))
		}

	}
}

//json
type Teacher struct {
	Name string			//首字母必须大写，不大写的无法输出
	Age int	`json:"no"`	//后面的json会把大写变成小写，或者重命名
	No int `json:"-"`	//不编码成json
	Phone int `json:",string"`		//转换成字符串再输出
}