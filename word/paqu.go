package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)
func main() {
	for i := 1;i <= 16; i++{
		var urlword = "http://www.etymon.cn/yingyucizhui/list_2_" + strconv.Itoa(i) + ".html"
		resp,err1 := http.Get(urlword)
		if err1 != nil {
			fmt.Println("http.Get err: ",err1)
		}
		//返回数据存放在Body中，通过读取Body获取相应的返回
		defer resp.Body.Close()

		buf := make([]byte,4*1024)

		//fmt.Println(resp.Status)	//返回码
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
		//fmt.Println(body)
		//将爬取到的数据写入到文件
		var filePath string
		filePath = "word_" + strconv.Itoa(i) + ".html"
		file,err3 := os.Create(filePath)
		if err3 != nil {
			fmt.Println("os.Create err: ",err3)
			return
		}
		defer file.Close()
		file.WriteString(body)
	}

}
