package main
import (
	"fmt"
)
type speaker interface{
	speak()
}
type cat struct{}
type dog struct{}
type person struct{}
func (c cat)speak(){
	fmt.Println("喵喵喵")
}
func (d dog)speak(){
	fmt.Println("汪汪汪")
}
func (p person)speak(){
	fmt.Println("啊啊啊啊")
}
func da(x speaker){
	//接收一个参数，传进来什么我就打什么
	x.speak()
}
func main() {
	var c1 cat
	var d1 dog
	da(c1)
	da(d1)
}