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
func f2(w http.ResponseWriter, r *http.Request) 
{
	fmt.Println(r.URL)
	fmt.Println(r.Method)
}
func main() {
	http.HandleFunc("/posts/Go/unit-test/", f1)
	http.HandleFunc("/hello/",f2)
	http.ListenAndServe("127.0.0.1:9090", nil)

}
