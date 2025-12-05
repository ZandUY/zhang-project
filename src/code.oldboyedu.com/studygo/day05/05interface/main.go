package main
import (
	"fmt"
)
type animal interface {
	move()
	eat(string)
}
type cat struct{
	name string
	feet int8
}
type chicken struct{
    feet int8
}
func (c chicken)move(){
	fmt.Println("鸡会飞")
}
func (c chicken)eat(food string){
	fmt.Printf("鸡吃%s",food)
}
func (c cat)move(){
	fmt.Println("猫会走")
}
func (c cat)eat(food string){
	fmt.Printf("猫吃%s\n",food)
}
func main() {
	var a1 animal
	bc:=cat{
		name:"小花",
		feet:4,
	}
	a1=bc
	fmt.Println(a1)
	a1.eat("方便面")
	fmt.Printf("%T\n",a1)
}