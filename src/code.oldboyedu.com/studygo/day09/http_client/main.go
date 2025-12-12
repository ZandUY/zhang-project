package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	// resp, err := http.Get("http://127.0.0.1:9090/hello/")
	// if err != nil {
	// 	fmt.Println("get url failed,err	:", err)
	// 	return
	// }
	data:=url.Values{}
	data.Set("name","张三")
	data.Set("age","18")
	queryStr:=data.Encode()
	urlObj,_:=url.Parse("http://127.0.0.1:9090/hello/")
	urlObj.RawQuery=queryStr
	req,_:=http.NewRequest("GET",urlObj.String(),nil)
	resp,_:=http.DefaultClient.Do(req)
	// resp.Body.Read()
	// resp.Body.Close()
	msg, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read resp body failed,err:", err)
		return
	}
	fmt.Println(string(msg))

}
