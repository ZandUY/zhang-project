/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
package main

import "fmt"

func yuanshuai(name string){
	fmt.Println("hello",name)
}
func low(f func()){
	f()
}
func f1(f func(string),name string)(func()){
	return func() {
		f(name)
	}
}
func main() {
	f4:=f1(yuanshuai,"理想")
	low(f4)
}
