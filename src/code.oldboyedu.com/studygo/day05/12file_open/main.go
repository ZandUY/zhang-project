package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFromFile1() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file err", err)
		return
	}
	for {
		var tmp = make([]byte, 128)
		n, err := fileObj.Read(tmp)
		if err != nil {
			fmt.Println("read file err", err)
			return
		}
		// fmt.Printf("读取了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			break
		}
	}

	defer fileObj.Close()
}
func readFromBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file err", err)
		return
	}
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("read file err", err)
			return
		}
		fmt.Print(line)
	}

}
func readFromIoutil() {
	ret, err := os.ReadFile("./main.go")
	if err != nil {
		fmt.Println("read file err", err)
		return
	}
	fmt.Println(string(ret))
}
func main() {
	//readFromFile1()
	readFromBufio()
	// readFromIoutil()
}
