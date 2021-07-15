package main

import (
	"fmt"
	"regexp"
)

func main() {
	for {
		fmt.Print("请输入: ")
		var str string
		fmt.Scan(&str)

		res := regexp.MustCompile(`hello(.*)`)
		resStr := res.FindAllStringSubmatch(str, -1)
		for _, v := range resStr {
			fmt.Println(v[1])
		}

	}
}
