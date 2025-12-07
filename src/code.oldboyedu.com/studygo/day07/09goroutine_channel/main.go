package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("机器%d 开始了作业:%d\n",id,j)
		results<-2*j
	}
	defer wg.Done()
}
var wg  sync.WaitGroup
func main() {
	jobs := make(chan int, 50)
	results := make(chan int,50)
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results)
	}
	for j:=1;j<=5;j++{
	    jobs<-j
	}
	close(jobs)
	// for res:=range results{
	// 	fmt.Println("结果:",res)
	// }
	 for i:=0;i<5;i++{
		fmt.Println(<-results)
	}
	// for res:=range results{
	// 	fmt.Println("结果:",res)
	// }
	wg.Wait()
	
}