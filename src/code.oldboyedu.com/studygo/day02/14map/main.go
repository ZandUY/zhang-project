package main

import (
	"fmt"
)

func main() {
	var m1 map[string]int
	fmt.Println(m1)
	m1=make(map[string]int,10)
	m1["理想"]=9000
	m1["jiwuming"]=35
	delete(m1,"理想")
	delete(m1,"不存在的key")
     for k,_:=range m1{
     	fmt.Println(k)
	 }
}
