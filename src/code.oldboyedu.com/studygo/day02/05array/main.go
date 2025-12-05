package main

import (
	"fmt"
)

func sum(a [6]int) int {
	sum := 0
	for _, i := range a {
		sum += i
	}
	return sum
}

func main() {
	// var a1 [3]bool
	//数组长度是数组类型的一部分
	// fmt.Println(a1)
	// a2 := [...]int{1, 2, 3, 4, 5}
	// fmt.Println(a2)
	// a3:=[5]int{0:1,4:2}
	// a4:=[...][2]int{{1,2},{3,4},{5,6}}
	// for _,i:=range a4{
	// 	for _,j:=range i{
	// 		fmt.Printf("%d\t",j)
	// 	}
	// 	fmt.Println()
	var a = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(sum(a))
	// for v, i := range a {
	// 	for j := v + 1; j < len(a); j++ {
	// 		if i+a[j] == 8 {
	// 			fmt.Printf("(%d,%d)\n", v, j)
	// 		}
	// 	}
	// }
}
