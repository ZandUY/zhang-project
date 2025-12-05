package main

import (
	"fmt"
)

func main() {
	// s1 := []int{1, 3, 5}
	// // fmt.Println(s1)
	// s2 := s1
	// s2[0]=100
	// fmt.Println(s1)
	// var s3 []int
	//  s3 = make([]int, 3, 5)
	// copy(s3, s1)
	// fmt.Println(s1, s2)
	// s2[0] = 100
	// fmt.Println(s1, s2)

	// s1=append(s1[:1],s1[2:]...)
	// fmt.Println(s1)

	// x1 := [...]int{1, 3, 5}
	// s1 := x1[:]
	// fmt.Println(s1, len(s1), cap(s1))
	// s1=append(x1[:1],9,9,9)
	// fmt.Println(s1, len(s1), cap(s1))
	// fmt.Println(x1)



	a1:=[...]int{1,3,5,7,9,11,13,15,17}
	s1:=a1[:]
	fmt.Println(s1)
	s2:=append(s1[:1],s1[2:]...)
	fmt.Println(s2)
	fmt.Println(a1)

}
