package main

import (
	"log"
	"os"
)

func main() {
	file,err:=os.OpenFile("log1.txt",os.O_CREATE|os.O_APPEND|os.O_RDWR,0644)
	if err!=nil{
		log.Fatalf("open log file error:%s\n",err)
		return 
	}
log.SetOutput(file)
	for {
		log.Println("这是一条测试日志")
	}
}
