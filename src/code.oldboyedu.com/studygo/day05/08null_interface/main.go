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
	var m1 map[string]interface{}
	m1=make(map[string]interface{},16)
	m1["name"]="沙河"
	m1["age"]=18
	m1["married"]=true
	m1["hobby"]=[...]string{"唱","跳","rap"}
	fmt.Println(m1)
	
}