package main
import(
	"fmt"
)
func f1(){
	fmt.Println("hello 沙河")
}
func f2() int{
	return 10
}
func f3(x func()){
	
}
func main(){
	a:=f1
	fmt.Printf("%T\n",a)
	b:=f2
	fmt.Printf("%T\n",b)
}