package main
import "fmt"
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("1", x, calc("10", x, y))
	x = 0
	defer calc("2", x, calc("20", x, y))
	y = 1
}