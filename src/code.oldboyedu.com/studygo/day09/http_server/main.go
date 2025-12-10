package main

import (
	"fmt"
	"net/http"
	"os"
)

func f1(w http.ResponseWriter, r *http.Request) {
	page, err := os.ReadFile("./page.html")
	if err != nil {
		fmt.Println("读取文件失败")
		return
	}

	w.Write(page)
}
func main() {
	http.HandleFunc("/posts/Go/unit-test/", f1)
	http.ListenAndServe("127.0.0.1:9090", nil)

}
