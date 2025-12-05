package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
}
func f1() {
	fmt.Println("测试github")
}
func f2() {
	fmt.Println("测试github2")
}
