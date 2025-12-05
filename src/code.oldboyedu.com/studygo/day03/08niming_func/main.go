package main
import(
	"fmt"
)
var f1=func (x,y int){
	fmt.Println(x+y)
}
func main(){
	func (x,y int){
		fmt.Println(x*y)
	}(10,20)
}