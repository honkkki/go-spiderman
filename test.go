package main

import (
	"fmt"
	"regexp"
)

func main()  {
	fmt.Println("请输入:")
	var str string;
	fmt.Scanln(&str)

	res := regexp.MustCompile(`hi.*`)

	//首页的文章标题链接
	resStr := res.FindString(str)

	fmt.Println(resStr)
}