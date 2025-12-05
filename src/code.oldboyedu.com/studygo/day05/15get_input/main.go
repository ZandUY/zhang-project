package main

import (
	"bufio"
	"fmt"
	"os"
)

func userScan() {
	var s string
	fmt.Printf("请输入内容")
	fmt.Scanln(&s)
	fmt.Printf("你输入的内容是：%s\n",s)
}
func userBufio(){
	var s string 
	fmt.Printf("请输入内容:")
	reader:=bufio.NewReader(os.Stdin)
	s,_=reader.ReadString('\n')
	fmt.Printf("您输入的内容是:%s\n",s)
}

func main() {
	userBufio()
}