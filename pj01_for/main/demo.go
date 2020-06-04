//Date: 2020-06-03
//Details:如何遍历字符串
package main

import "fmt"

func main() {

	//字符串遍历方式1-传统方式
	var str string = "hello,world洪荒"
	str2 := []rune(str) //就是把str转成[]rune
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i]) //下标...
	}

	//字符串遍历方式2-for-range
	for index, val := range str {
		fmt.Printf("index = %d,val = %c\n", index, val)
	}
}
