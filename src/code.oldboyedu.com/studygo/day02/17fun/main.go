package main

import (
	"fmt"
)
func add(x  int,y int) int{
	return x+y
}
func f2(){
	fmt.Println("woaini")
}

func f3() (ret int){//命名返回值
	ret=100
	return 
}

func  f4()(x int ,y int){
	return 1,2
}
func f5(x,y,z int){
	fmt.Println(x,y,z)
}
func f7(x int ,y ...int){
	fmt.Println(y)
	//y的类型是一个切片
	//可变长参数必须放在函数参数的最后
}
func main() {
	r:=add(1,2)
	fmt.Println(r)
	f2()
	fmt.Println(f3())
	a,b:=f4()
	fmt.Println(a,b)
	f7(1,2,3,4,5)
}
