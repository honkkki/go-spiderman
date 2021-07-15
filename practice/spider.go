package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//https://www.pengfu.com/xiaohua_1.html
//https://www.pengfu.com/xiaohua_2.html


func SpiderPage(i int, ch chan int)  {
	//url
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取第%d页的数据:%s\n", i, url)

	result, err := HttpGet(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println("爬取到的网页内容:", result)
	res := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if res == nil {
		return
	}
	//首页的文章标题链接
	joyUrl := res.FindAllStringSubmatch(result, -1)
	fileTitle := make([]string, 0)
	fileContent := make([]string, 0)



	//fmt.Println(joyUrl)
	for _, data := range joyUrl {
		title, content, err1 := SpiderJoy(data[1])
		if err1 != nil {
			fmt.Println(err1)
			continue
		}

		fileTitle = append(fileTitle, title)
		fileContent = append(fileContent, content)
	}
	//fmt.Println(fileTitle)
	JoyToFile(i, fileTitle, fileContent)
	ch <- i
}

func JoyToFile(i int, title, content []string)  {
	file, _ := os.Create(strconv.Itoa(i)+".txt")
	defer file.Close()
	//
	n := len(title)
	for i := 0; i < n; i++ {
		file.WriteString(title[i]+"\n")
		file.WriteString(content[i]+"\n")
		file.WriteString("===========================\n")
	}

}


func SpiderJoy(url string) (title , content string, err error) {
	result, _ := HttpGet(url)
	res := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if res == nil {
		err = errors.New("compile error")
		return
	}
	//get
	tmpTitle := res.FindAllStringSubmatch(result, 1)
	for _, data := range tmpTitle {
		title = data[1]
		title = strings.Replace(title, "\t", "", -1)
		break
	}

	contentRes := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	if contentRes == nil {
		err = errors.New("compile error")
		return
	}
	//get
	tmpContent := contentRes.FindAllStringSubmatch(result, 1)
	for _, data := range tmpContent {
		content = data[1]
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "&nbsp;", "", -1)
		content = strings.Replace(content, " ", "", -1)
		content = strings.Replace(content, "<br />", "", -1)
		content = strings.Replace(content, "<br/>", "", -1)
		break
	}
	return
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
	fmt.Println("spidering")

	ch := make(chan int)
	for i := start; i <= end; i++ {
		//定义函数爬主页面
		go SpiderPage(i, ch)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-ch)
	}

}

func main()  {
	DoWork(1, 5)
}
