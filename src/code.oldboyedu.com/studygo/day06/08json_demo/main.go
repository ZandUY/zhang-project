package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"理想","age":18}`
	var p person
	json.Unmarshal([]byte(str),&p)
	fmt.Println(p.Name,p.Age)
}