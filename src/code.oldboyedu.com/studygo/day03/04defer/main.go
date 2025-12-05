package main
import(
	"fmt"
)
func deferDemo(){
	fmt.Println(123)
	defer fmt.Println(456)
	fmt.Println(789)
	fmt.Println("101112")
	fmt.Println("131415")
}
func main(){
	deferDemo()
}