package main

import (
	"fmt"
)

func main() {
	n := 18
	fmt.Println(&n)
	p := &n
	fmt.Printf("%T\n", p)
	fmt.Println(*p)
	var p2 *int = p
	fmt.Println(*p2)
}
