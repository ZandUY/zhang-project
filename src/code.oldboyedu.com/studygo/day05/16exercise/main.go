package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFromFile1() {
	fileObj, err := os.OpenFile("./ex.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("open file error:%s\n", err)
		return
	}
	defer fileObj.Close()
	for {
		var tmp = make([]byte, 128)
		n, err := fileObj.Read(tmp)
		if err != nil {
			fmt.Printf("read file error:%s\n", err)
			return
		}
		fmt.Printf("读取了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}

}
func readFromBufio() {
	fileObj, err := os.OpenFile("./ex.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("open file error:%s\n", err)
	}
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	for {

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read file error:%s\n", err)
			return
		}
		fmt.Print(line)
	}
}

func readFromIoutil() {
	c, err := os.ReadFile("./ex.txt")
	if err != nil {
		fmt.Printf("read file error:%s\n", err)
		return
	}
	fmt.Println(string(c))
}
func writeToFile1() {
	fileObj, err := os.OpenFile("./ex.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file error:%s\n", err)
		return
	}
	defer fileObj.Close()
	fileObj.WriteString("时运不齐\n")
	fmt.Fprintln(fileObj, "时运不齐2")
}
func writeToFile2() {
	fileObj, err := os.OpenFile("./ex.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file error:%s\n", err)
		return
	}
	defer fileObj.Close()
	writer := bufio.NewWriter(fileObj)
	writer.WriteString("冯\n")
	writer.Flush()
}
func writeToFile3() {
	// fileObj,err:=os.OpenFile("./ex.txt", os.O_WRONLY|os.O_APPEND, 0644)
	str := "语"
	err:= os.WriteFile("./ex.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("write file error:%s\n", err)
		return
	}
}
func writeToFile4(){
	var str string
	fmt.Printf("请输入内容:")
	reader:=bufio.NewReader(os.Stdin)
	str,_=reader.ReadString('\n')
	err:= os.WriteFile("./ex.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("write file error:%s\n", err)
		return
	}
}
func main() {
	// readFromFile1()
	// readFromBufio()
	// readFromIoutil()
	// writeToFile1()
	// writeToFile2()
	// writeToFile3()
	writeToFile4()
}
