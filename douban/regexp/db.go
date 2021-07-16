package main

import (
	"fmt"
	"github.com/go-spiderman/utils"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func spider1(url string)  {
	defer wg.Done()
	body := utils.Fetch(url)
	body = strings.Replace(body, "\n", "", -1)
	re := regexp.MustCompile(`<div class="hd">(.*?)</div>`)
	titleRe := regexp.MustCompile(`<span class="title">(.*?)</span>`)
	idRe := regexp.MustCompile(`<a href="https://movie.douban.com/subject/(\d+)/"`)
	items := re.FindAllStringSubmatch(body, -1)

	// div hd下套着title id
	for _, item := range items {
		fmt.Println(titleRe.FindStringSubmatch(item[1])[1], idRe.FindStringSubmatch(item[1])[1])
	}
}

func main()  {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go spider1("https://movie.douban.com/top250?start=" + strconv.Itoa(i*25) + "&filter=")
	}

	wg.Wait()
	used := time.Since(start)
	fmt.Println("used time:", used)
}
