package main

import (
	"fmt"
	// "strings"
	// "unicode"
)

func main() {
	// s := "我爱你123456sfjsk冯f"
	// res := 0
	// for _, v := range s {
	// 	if unicode.Is(unicode.Han, v) {
	// 		res++
	// 	}
	// }
	// fmt.Println("汉字有：", res)

	//2. how do you do
	// s := "how do you do"
	// words:=strings.Split(s," ")
	// fmt.Println(words)
	// var m=make(map[string]int,50)
	// for _,v :=range words{
	// 	m[v]++
	// }
	// fmt.Println(m)


	//回文字符串
	s := "上海自来水来自海上"
	s1:=[]rune(s)
	for i,v :=range s{
		fmt.Printf("index=%d,value=%c\n",i,v)
	}
	for i,v :=range s1{
		fmt.Printf("index=%d,value=%c\n",i,v)
	}
	flag:=true
	for i,j:=0,len(s1)-1;i<j;i,j=i+1,j-1{
		if s1[i]!=s1[j]{
			fmt.Println("不是回文字符串")
			flag=false
			break
		}
	}
	if flag{
		fmt.Println("是回文字符串")
	}


	

}
