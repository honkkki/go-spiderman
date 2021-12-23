package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/honkkki/go-spiderman/utils"
	"log"
	"strconv"
	"strings"
	"time"
)

var banner = `

.-----..-----..-----.
|  _  ||__ --||  _  |
|___  ||_____||   __|
|_____|       |__|

🍭🍭🍭🍭🍭🍭🍭🍭🍭🍭🍭🍭🍭🍭🍭🍭🍭

`

func spinner(delay time.Duration) {
	finish := make(chan struct{})

	go func() {
		worker()
		finish<- struct{}{}
	}()

	for {
		select {
		case <-finish:
			fmt.Println("finish!")
			return

		default:
			utils.Spinner(delay)
		}
	}
}

// worker 耗时操作
func worker() {
	doc := utils.FetchQuery("https://www.v2ex.com")
	if doc == nil {
		log.Fatal("oops! fetch url fail...")
	}

	doc.Find("#Rightbar").Find("#TopicsHot").Find(".cell").Each(func(i int, selection *goquery.Selection) {
		if i!=0 {
			fmt.Println()
			fmt.Println("#" + strconv.Itoa(i))
			title := selection.Find(".item_hot_topic_title").Text()
			title = strings.TrimSpace(title)
			fmt.Println(title)
			fmt.Println()
			url, _ := selection.Find(".item_hot_topic_title").Find("a").Attr("href")
			fmt.Println("https://www.v2ex.com" + strings.TrimSpace(url))
			fmt.Println("------------------------------------")
		}
	})
}

func main() {
	//banner, _ := ioutil.ReadFile("../resource/banner.txt")
	fmt.Printf(banner)
	spinner(100 * time.Millisecond)
}
