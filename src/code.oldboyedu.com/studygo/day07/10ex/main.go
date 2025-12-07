package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

var wg sync.WaitGroup

func genInt64(jobChan chan<- int64) {
	defer wg.Done()
	for{
		jobChan <- rand.Int64()%100
		time.Sleep(time.Millisecond*500)
	}
}
func sumOfNum(id int,jobChan <-chan int64, resultChan chan<- int64) {
	defer wg.Done()
	for {
		num:=<-jobChan
		var sum int64 = 0
		f:=num
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		fmt.Printf("机器:%d 随机数:%d 各位数字之和为:  %d\n", id,f,sum)
		resultChan <- sum
	}
}
func main() {
	jobChan := make(chan int64,50)
	resultChan := make(chan int64,50)
	wg.Add(1)
	go genInt64(jobChan)
	for i := 0; i < 24; i++ {
		wg.Add(1)
		go sumOfNum(i,jobChan, resultChan)
	}
	wg.Wait()
}
