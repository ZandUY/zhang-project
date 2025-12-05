package main
import (
	"fmt"
)
type car interface{
	run()
}
type falali struct {
    brand string
}
type baoshijie struct {
    brand string
}
func drive(c car){
	c.run()
}
func (f falali)run(){
	fmt.Println(f.brand,"开车了")
}
func (b baoshijie)run(){
	fmt.Println(b.brand,"开车了")
}
func main() {
	var f1=falali{
		brand: "法拉利",
	}
	var b1=baoshijie{
		brand: "保时捷",
	}
	drive(f1)
	drive(b1)
}