package main

import (
	"bytes"
	"fmt"
	"regexp" //正则表达式依赖包
)

func main() {
	fmt.Println("正则表达式")
	flag1,err1 := regexp.Compile(`a.b`)	//Compile用于检测正则表达式是否合法
	if err1 == nil {
		fmt.Println("flag1:",flag1)		//a.b
		//返回参数*Regexp（带有Compile的函数返回的）拥有查找功能
		fmt.Println("flag1.FindString:",flag1.FindString("aba bed acb adb"))	//FindString返回匹配到的第一个
		fmt.Println("flag1.Find:",flag1.Find([]byte("aba bed acb adb")))	//[97 32 98]
		fmt.Println("flag1.FindAll:",flag1.FindAll([]byte("aba bed acb adb"),10))	//[[97 32 98] [97 99 98]],只查找前10个匹配
		fmt.Println("flag1.Find:",flag1.FindAllString("aba bed acb adb",10))
	}else {
		fmt.Println(err1)
	}

	flag2,err2 := regexp.MatchString(`a.b`,"aba bed acb adb")	//MatchString用于检测后面字符串中时候包含正则表达式所匹配内容
	if err2 == nil {
		fmt.Println("flag2:",flag2)		//true
	}else {
		fmt.Println(err2)
	}

	flag3,err3 := regexp.Match(`a.b`,[]byte("aba bed acb adb"))	//Match用于检测后面字节数组中时候包含正则表达式所匹配内容
	if err3 == nil {
		fmt.Println("flag3:",flag3)		//true
	}else {
		fmt.Println(err3)
	}

	flag4,err4 := regexp.CompilePOSIX(`a.b`)
	// 与Compile不同的是，CompilePOSIX 使用 POSIX 语法，
	// 同时，它采用最左最长方式搜索，
	// 而 Compile 采用最左最短方式搜索
	// POSIX 语法不支持 Perl 的语法格式：\d、\D、\s、\S、\w、\W
	if err4 == nil {
		fmt.Println("flag4:",flag4)		//a.b
	}else {
		fmt.Println(err4)
	}

	r := bytes.NewReader([]byte("Hello World!"))
	flag5,err5 := regexp.MatchReader(`H*`,r) //判断在 r 中能否找到正则表达式 pattern 所匹配的子串
	if err5 == nil {
		fmt.Println("flag5:",flag5)		//true
	}else {
		fmt.Println(err5)
	}

	flag6 := regexp.MustCompile(`a.b`)
	// 与Compile不同的是，当正则表达式 str 不合法时，MustCompile 会抛出异常
	// 而 Compile 仅返回一个 error 值
	fmt.Println("flag6:",flag6)		//a.b

	fmt.Println(regexp.QuoteMeta(`\.+*?()|[]{}^$`))		//将字符串 s 中的“特殊字符”转换为其“转义格式“
	// 特殊字符有：\.+*?()|[]{}^$
	// 这些字符用于实现正则语法，所以当作普通字符使用时需要转换
	//可以用于查找这个字符串的转移字符
}
