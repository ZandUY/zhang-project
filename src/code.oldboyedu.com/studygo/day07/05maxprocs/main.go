package main

import (
	"fmt"
	"runtime"
	"sync"
)

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}
func b(){
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}

}
var wg sync.WaitGroup
func main(){
	runtime.GOMAXPROCS(1)
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}