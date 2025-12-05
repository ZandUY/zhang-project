package main

import (
	"fmt"
)

func main() {
	s := "Hello, 世界"
	n := len(s)
	fmt.Println("字符串长度：", n)
	// for i := 0; i < n; i++{
	// 	// fmt.Println(s[i])
	// 	fmt.Printf("%c\n", s[i])
	// }
	for _,c:=range s{
		fmt.Printf("%c\n", c)

	}

	var flag bool=true
	fmt.Println(flag)
}