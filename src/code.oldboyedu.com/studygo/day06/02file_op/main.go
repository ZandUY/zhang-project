package main

import (
	"fmt"
	"os"
)

func f1() {
	fileObj, err := os.OpenFile("./sb.txt", os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file error:%s\n", err)
		return
	}
	defer fileObj.Close()
	_,err1:=fileObj.Seek(2,0)
	if err1!=nil{
		fmt.Printf("seek file error:%s\n", err)
		return
	}
	_,err2:= fileObj.Write([]byte("\nf\n"))
	if err2!=nil{
		fmt.Printf("write file error:%s\n", err)
		return
}
}
func main() {
	f1()
}