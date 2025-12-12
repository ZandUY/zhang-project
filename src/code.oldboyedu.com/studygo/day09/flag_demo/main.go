package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "张三", "请输入姓名")
	flag.Parse()
	fmt.Println(*name)
}