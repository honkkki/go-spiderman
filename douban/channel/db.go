package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-spiderman/utils"
	"log"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type Job struct {
	Id  int
	Url string
}

type Res struct {
	Id    string
	title string
}

type jobChan chan Job
type resChan chan Res

var (
	mu   sync.Mutex
	num  int
	done = make(chan struct{})
	wg sync.WaitGroup
)

func spider(url string, rc resChan) {
	doc := utils.FetchQuery(url)
	if doc == nil {
		log.Fatal("初始化doc失败")
	}

	doc.Find("ol.grid_view li").Find(".hd").Each(func(i int, selection *goquery.Selection) {
		imgUrl, _ := selection.Find("a").Attr("href")
		title := selection.Find(".title").Eq(0).Text()
		id := strings.Split(imgUrl, "/")[4]
		rc <- Res{
			id,
			title,
		}
		mu.Lock()
		num++
		if num == 250 {
			close(rc)
		}
		mu.Unlock()

	})
}

func worker(jc jobChan, rc resChan) {
	defer wg.Done()
	for job := range jc {
		fmt.Println("working id:", job.Id)
		spider(job.Url, rc)
	}
}

func run(jc jobChan, rc resChan) {
	for i := 0; i < 5; i++ {
		go worker(jc, rc)
	}
}

func printRes(rc resChan) {
	for res := range rc {
		fmt.Println(res.Id, res.title)
	}
	done<- struct{}{}
}

func main() {
	jc := make(jobChan, 10)
	rc := make(resChan, 250)
	//start := time.Now()

	wg.Add(5)
	run(jc, rc)
	go printRes(rc)
	for i := 0; i < 10; i++ {
		url := "https://movie.douban.com/top250?start=" + strconv.Itoa(i*25) + "&filter="
		job := Job{
			Id:  i + 1,
			Url: url,
		}

		jc <- job
	}

	fmt.Println(runtime.NumGoroutine())
	close(jc)
	wg.Wait()

	// 等待打印完毕
	<-done
	fmt.Println(num)
	//used := time.Since(start)
	//fmt.Println("used time:", used)
}
