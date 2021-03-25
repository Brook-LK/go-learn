package main

//爬取词根词缀
//url:http://www.etymon/yingyucizhui/list_2_1.html

//获取深层次的讲解，有两种格式开头
//1:<dt class="highlighth"><a href="   地址    结尾     ">  爬取标题  </a>
//2:<dt class=""><a href="   地址    结尾     ">  爬取标题  </a>

//爬取内容
//1:</dt><dd class="highlighth">   爬取内容   </dd>
//2:</dt><dd class="">   爬取内容   </dd>

//爬取详细内容，进入地址后
//</dt><dd class="highlight">    详细内容    </dl>

import (
	"fmt"
	"io"
	"net/http"
)
func main() {
	resp,err1 := http.Get("http://www.etymon.cn/yingyucizhui/list_2_1.html")	//list_2_ + 1到16 + .html
	//resp,err1 := http.Get("http://www.etymon.cn/yingyucizhui/yingyuqianzhui/43009.html")
	if err1 != nil {
		fmt.Println("http.Get err: ",err1)
	}

	defer resp.Body.Close()

	buf := make([]byte,4*1024)

	//fmt.Println(resp.Status)
	//fmt.Println(resp.StatusCode)
	//fmt.Println(resp.Header)
	//fmt.Println(resp.Body)
	var body string
	for {
		n,err2 := resp.Body.Read(buf)
		body += string(buf[:n])		//不管爬取成功与否，先将爬取的数据些进去，如果先break后面的就写不进去了
		if err2 != nil{
			if err2 == io.EOF {
				fmt.Println("爬取成功")
			}else{
				fmt.Println("resp.Body.Read err: ",err2)
			}
			break
		}
	}
	fmt.Println(body)

}
