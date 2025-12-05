package main

import "fmt"

//111
//左边的1表示吃饭
//中间的1表示睡觉
//右边的1表示打豆豆
const(
	eat int=4
	sleep int=2
	da int =1
)
func f(arg int){
	fmt.Printf("%b\n",arg)
}
func main(){
	f(eat|da)
}