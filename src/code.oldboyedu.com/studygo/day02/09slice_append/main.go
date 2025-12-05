package main
import (
	"fmt"
)

func main() {
	s1:=[]string{"北京","上海","深圳"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n",s1,len(s1),cap(s1) )
	s1=append(s1,"广州","杭州","成都")
		fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n",s1,len(s1),cap(s1) )
	fmt.Println(s1)
	// fmt.Println(s1...)
}