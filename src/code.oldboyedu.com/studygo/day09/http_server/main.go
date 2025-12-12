package main

import (
	"fmt"
	"io"
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
func f2(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	fmt.Println(r.URL)
	fmt.Println(queryParams.Get("name"), queryParams.Get("age"))
	fmt.Println(r.Method)
	fmt.Println(io.ReadAll(r.Body))
	w.Write([]byte("ok"))
}
func main() {
	http.HandleFunc("/posts/Go/unit-test/", f1)
	http.HandleFunc("/hello/", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)

}
