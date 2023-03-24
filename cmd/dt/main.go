package main

import (
	"log"

	"github.com/honkkki/go-spiderman/internal/app/dt"
)

func main() {
	s := dt.NewSpider("dt-SpiderMan")
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}

	log.Println(s.AppName(), "job finish!")
}
