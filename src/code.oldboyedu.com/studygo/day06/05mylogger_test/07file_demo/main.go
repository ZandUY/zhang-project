package main

import (
	"fmt"
	"os"
)

func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}
	fmt.Printf("%T\n",fileObj)
	fileInfo,err:=fileObj.Stat()
	if err!=nil{
		fmt.Println("get file info failed", err)
		return
	}
	fmt.Printf("文件大小是:%dB",fileInfo.Size())
	fmt.Printf("文件名是:%s\n",fileInfo.Name())
}