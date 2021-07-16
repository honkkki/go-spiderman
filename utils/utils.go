package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) string {
	fmt.Println("fetching url:", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("request url fail:", err)
		return ""
	}

	defer resp.Body.Close()
	if code := resp.StatusCode; code != 200 {
		fmt.Println("http code error:", code)
		return ""
	}

	data, _ := ioutil.ReadAll(resp.Body)
	return string(data)
}

func FetchQuery(url string) *goquery.Document {
	fmt.Println("fetching url:", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("request url fail:", err)
		return nil
	}

	defer resp.Body.Close()
	if code := resp.StatusCode; code != 200 {
		fmt.Println("http code error:", code)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("new doc fail:", err)
		return nil
	}

	return doc
}
