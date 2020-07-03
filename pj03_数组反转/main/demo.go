//Date : 2020-07-03
//Details: 随机生成5个数，并将其反转打印
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//思路
	//1.随机生成5个数，rand.Intn()
	//2.得到随机数后，就放到一个数组中 int数组
	//3.将其反转打印,交换的次数是 len / 2,倒数第一个元素和第一个元素交换，倒数第二个和第二个交换
	var intArr3 [5]int
	//为了每次生成的随机数不一样，我们需要给一个seed值
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr3); i++ {
		intArr3[i] = rand.Intn(100)
	}
	fmt.Println("交换前：", intArr3)
	temp := 0
	for i := 0; i < len(intArr3)/2; i++ {
		temp = intArr3[len(intArr3)-1-i]
		intArr3[len(intArr3)-1-i] = intArr3[i]
		intArr3[i] = temp
	}
	fmt.Println("交换后：", intArr3)
}
