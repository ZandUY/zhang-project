package calca

import "fmt"

func Add(x, y int) int {
	return x + y
}
func init() {
	fmt.Println("calca 包被引用了")
}