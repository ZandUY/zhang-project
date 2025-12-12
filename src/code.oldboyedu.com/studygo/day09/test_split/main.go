package main

import (
	"fmt"

	splitstring "code.oldboyedu.com/studygo/day09/split_string"
)

func main() {
	ret:=splitstring.Split("babcbef","b")
	fmt.Printf("%#v\n",ret)
}