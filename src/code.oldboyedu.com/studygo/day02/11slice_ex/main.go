package main

import (
	"fmt"
	"sort"
)

func main() {
	var a1=[5]int{3,7,8,9,1}
	sort.Ints(a1[:])
	fmt.Println(a1)
}