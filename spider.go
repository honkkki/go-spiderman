package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//https://www.pengfu.com/xiaohua_1.html
//https://www.pengfu.com/xiaohua_2.html


func SpiderPage(i int)  {
	//url
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取第%d页的数据:%s\n", i, url)

	result, err := HttpGet(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("爬取到的网页内容:", result)


}

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 1024 * 4)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}

		result += string(buf[:n])
	}
	return
}


func DoWork(start, end int)  {
	fmt.Println("spider ing")

	for i := start; i <= end; i++ {
		//定义函数爬主页面
		SpiderPage(i)
	}

}

func main()  {

	DoWork(1, 2)
}
