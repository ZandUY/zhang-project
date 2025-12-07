package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func f1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}
func f2(ch1 ,ch2 chan int) {
	defer wg.Done()

	for x:=range ch1{
		ch2<-x*x
	}
	close(ch2)
}
func main() {
	a:=make(chan int,100)
	b:=make(chan int,100)
	wg.Add(2)
	go f1(a)
	go f2(a,b)
	wg.Wait()
for x:=range b{
	fmt.Println(x)
}

}
