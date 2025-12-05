package main

// import (
// 	"fmt"
// "..\calca\calc"
// )
import (
	"fmt"

	"code.oldboyedu.com/studygo/day05/calca"
)

// 主函数入口
func main() {
	// 调用calc包中的Add函数，传入参数1和2，并将结果赋值给ret变量
	ret := calca.Add(1, 2)
	// 打印ret变量的值
	fmt.Println(ret)
}
