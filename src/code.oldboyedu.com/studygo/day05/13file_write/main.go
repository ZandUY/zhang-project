package main

import (
	"bufio"
	"fmt"
	"os"
)
func writeDemo2() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	defer fileObj.Close()
	wr:=bufio.NewWriter(fileObj)
	wr.WriteString("hello 沙河\n")
	wr.Flush()
}

func writeDemo3(){
		str :="我爱你北京"
		err:=os.WriteFile("./xx.txt", []byte(str), 0666)
		if err !=nil{
			fmt.Printf("write file failed,err:%v\n", err)
				return
		}
}
// 主函数入口
func main() {
// writeDemo2()
writeDemo3()
}
