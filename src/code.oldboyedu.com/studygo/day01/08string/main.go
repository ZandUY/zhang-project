package main

//字符串
import (
	"fmt"
	"strings"
)

func main() {
	path := "D:\\Go\\src\\code.oldboyedu.com\\studygo\\day01"
	// fmt.Println(path)
	// s2:=`南山南
	// 北秋北
	// 每个人都有自己的苦衷`
	// fmt.Println(s2)
	name := "理想"
	word := "dsb"
	fmt.Println(name + word)
	ss := name + word
	fmt.Println(ss)
	//返回新字符串
	s1 := fmt.Sprintf("%s%s", name, word)
	fmt.Println(s1)
	//分割
	res := strings.Split(path, "\\")
	fmt.Println(res)
	fmt.Println(len(word))
	fmt.Println(strings.Contains(name,"理性"))
	fmt.Println(strings.Contains(name,"理想"))

	fmt.Println(strings.HasPrefix(name,"理"))
	fmt.Println(strings.HasSuffix(name,"想"))

	s2:="abcdefgh"
	fmt.Println(strings.Index(s2,"cde"))

	fmt.Println(strings.Join(res,"+"))

}
