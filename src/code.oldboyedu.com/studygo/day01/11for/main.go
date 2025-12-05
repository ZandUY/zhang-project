package main

import (
	"fmt"
)

func main() {
	// //1.基本
	// for i:=0;i<10;i++{
	// 	fmt.Println(i)
	// }
	// //2.变种1
	// var i=5
	// for ;i<10;i++{
	// 	fmt.Println(i)
	// }
	//3.变种2
	// s := "Hello沙河"
	// for i, v := range s {
	// 	fmt.Printf("索引:%d,字符:%c\n", i, v)
	// }
	for i:=1;i<=9;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%d*%d=%d ",j,i,j*i)
	}
		fmt.Println()
}
}
