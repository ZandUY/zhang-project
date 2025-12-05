package main

import (
	"fmt"
)

func main() {
	// var s1 = make([]map[string]int, 10, 10)
	// s1[0] = make(map[string]int, 2)
	// s1[0]["b"] = 100
	// s1[0]["a"] = 200
	// fmt.Println(s1)
	var m1=make(map[string][]int,10)
	m1["北京"]=[]int{10,20,30}
	fmt.Println(m1)
}
