package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// sec := now.Unix()
	// ret := time.Unix(int64(sec), 0)
	// fmt.Println(ret)
	// fmt.Println(ret.Add(24*time.Hour))
	// timer:=time.Tick(time.Second)
	// for t := range timer{
	// 	fmt.Println(t)
	// }
	// fmt.Println(now.Format("2006-01-02"))
	// fmt.Println(now.Format("2006/01/02 15/04/05.000 AM"))
	nextYear, err := time.Parse("2006-01-02 15:04:05", "2025-12-01 17:08:00")
	if err != nil {
		fmt.Printf("parse time error:%s\n", err)
		return
	}
	nextYear = nextYear.UTC()
	now = now.UTC()
	d := now.Sub(nextYear)
	fmt.Println(d)
	// fmt.Println("开始sleep了")
	// time.Sleep(5 * time.Second)
	// fmt.Println("5秒钟过去了")
	// f2()
}
func f1() {

}
func f2() {
	now := time.Now()
	fmt.Println(now)
	// time.Parse("2006-01-02 15:04:05", "2025-12-01 19:01:00")
	locObj, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load location error:%s\n", err)
		return
	}
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2025-12-01 19:12:00", locObj)
	if err != nil {
		fmt.Printf("parse time error:%s\n", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))

}
