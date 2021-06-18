package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

//https://www.douyu.com/g_yz


//读取网页源代码
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

func saveImg(id int, url string, ch chan int)  {
	path := "./images/" + strconv.Itoa(id+1) + ".jpg"
	f, _ := os.Create(path)
	defer f.Close()

	resp, err1 := http.Get(url)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 1024 * 4)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}

		f.Write(buf[:n])
	}
	ch <- id+1
	return
}


func main()  {
	url := "https://www.douyu.com/g_yz"
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	res := regexp.MustCompile(`<img loading="lazy" src="(?s:(.*?))"`)

	all := res.FindAllStringSubmatch(result, 10)


	ch := make(chan int)

	for id, imgUrl := range all {
		fmt.Println(id, imgUrl[1])
		go saveImg(id, imgUrl[1], ch)
	}

	n := len(all)
	for i := 0; i < n; i++ {
		fmt.Printf("第%d张图片下载完成\n", <-ch)
	}

}
