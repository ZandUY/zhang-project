package main
import (
	"fmt"
)

type animal interface{
    mover
    eater
}
type mover interface {
	move()
}
type eater interface {
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

}