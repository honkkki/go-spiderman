package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/honkkki/go-spiderman/model"
	"github.com/honkkki/go-spiderman/utils"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func spider2(url string) {
	defer wg2.Done()
	doc := utils.FetchQuery(url)
	if doc == nil {
		log.Fatal("初始化doc失败")
	}

	doc.Find("ol.grid_view li").Find(".hd").Each(func(i int, selection *goquery.Selection) {
		imgUrl, _ := selection.Find("a").Attr("href")
		title := selection.Find(".title").Eq(0).Text()
		id := strings.Split(imgUrl, "/")[4]

		// save to db
		data := &model.Movie{
			MovieNo: id,
			Title:   title,
		}

		model.CreateMovie(data)
		//fmt.Println(title, id)
	})
}

func main() {
	model.InitDB("../../config.ini")
	model.DB.Exec("TRUNCATE TABLE movie;")
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go spider2("https://movie.douban.com/top250?start=" + strconv.Itoa(i*25) + "&filter=")
	}

	wg2.Wait()
	used := time.Since(start)
	fmt.Println("used time:", used)
}
