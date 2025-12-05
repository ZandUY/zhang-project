package main

import (
	"fmt"
)

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}
func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}
func f3(f func(int, int), x, y int) func() {
	return func() {
		f(x, y)
	}
}
func main() {
	ret := f3(f2, 1, 3)
	f1(ret)
}
