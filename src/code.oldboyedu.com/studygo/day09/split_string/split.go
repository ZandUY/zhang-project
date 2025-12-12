package splitstring

import "strings"

func Split(str string, sep string) []string {
	var ret []string
	index := strings.Index(str, sep)

	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	ret = append(ret, str)
	return ret
}
func Fibo(n int) int {
	if n==1 || n==2 {
	    return 1
	}
	return Fibo(n-1)+Fibo(n-2)
}
