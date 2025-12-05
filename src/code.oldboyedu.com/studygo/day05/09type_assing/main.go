package main
import (
	"fmt"
)
func assign(a interface{}){
	switch a.(type){
	case string:
		fmt.Println("string",)
	case int:
		fmt.Println("int",a)
	case bool :
		fmt.Println("bool",a)
	default:
		fmt.Println("unknown")
	}
}
func main() {
	assign(100)
	assign(int64(100))
	assign("哈哈哈")
}