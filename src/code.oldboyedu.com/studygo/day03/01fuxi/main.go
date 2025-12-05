package main

import "fmt"

//day03复习
func main(){
	// var ages [30]int
	// fmt.Println(ages)
	// var ages2 [30]int
	// ages2=[30]int{1,2,3,4,5}
	// fmt.Println(ages2)
	s:="12346"
	var p *string=&s
	fmt.Println(*p)
	p=0x123
}