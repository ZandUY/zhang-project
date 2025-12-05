package main

import "fmt"

type dog struct {
	name string
}

func newDog(n string) *dog {
	return &dog{
		name: n,
	}
}
func (d dog) wang() {
	fmt.Println(d.name, "汪汪汪")
	d.name="大旺财"
	fmt.Println(d.name)
}
func main() {
	d1 := newDog("旺财")
	fmt.Println(d1)
	d1.wang()
	fmt.Println(d1)
}
