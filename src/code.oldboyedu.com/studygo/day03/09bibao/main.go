package main

import (
	"fmt"
)

func f1(f func()) {
	f()
}
func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}
func f3(f func(int,int), x, y int) func() {
		tmp := func() {
			f(x, y)
		}
		return tmp
	
}

// 主函数，程序的入口点
func main() {
	// 调用f1函数，并将f2函数作为参数传递
	ret := f3(f2, 100, 200)
	f1(ret)
}
