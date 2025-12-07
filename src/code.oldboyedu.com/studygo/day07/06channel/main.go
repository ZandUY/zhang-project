package main

import "fmt"

var ch1 chan int
var a []int

func main() {
	ch1 =make(chan int)
	fmt.Println(ch1)
	
}