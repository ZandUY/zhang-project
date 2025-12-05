package main

import "fmt"
type person struct{
	name string
	age int 
}
func newPerson(n string ,a int)*person{
	return &person{
		name:n,
		age:a,
	}
}
func main() {
	p1:=newPerson("张三",20)
	// fmt.Println(p1)
	fmt.Println(p1.name,p1.age)
}