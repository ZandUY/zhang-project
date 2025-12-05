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
func (c *chicken)move(){
	fmt.Println("鸡会飞")
}
func (c *chicken)eat(food string){
	fmt.Printf("鸡吃%s",food)
}
func (c *cat)move(){
	fmt.Println("猫会走")
}
func (c *cat)eat(food string){
	fmt.Printf("猫吃%s\n",food)
}
func main() {
	var a1 animal
	c1:=cat{"tom",4}
	c2:=&cat{"kitty",4}
	a1=&c1
	a1=c2
	fmt.Println(a1)
}