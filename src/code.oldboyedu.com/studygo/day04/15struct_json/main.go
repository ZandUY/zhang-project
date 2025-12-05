package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name`
	Age  int  `json:"age"`
}

func main() {
	// p1 := person{
	// 	Name: "沙河娜扎",
	// 	Age:  18,
	// }
	// b, err := json.Marshal(p1)
	// if err != nil {
	// 	fmt.Printf("json marshal failed,err:%v", err)
	// 	return
	// }
	// fmt.Printf("%v",string(b))
	str:=`{"name":"理想","age":18}`
	var p2 person
	json.Unmarshal([]byte(str),&p2)
	fmt.Printf("%#v\n",p2)

}
