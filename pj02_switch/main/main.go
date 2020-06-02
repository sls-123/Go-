//Date: 2020-05-30
//Details: 请编写一个程序，该程序可以接受一个字符，比如：a,b,c,d,e,f,g a表示星期一，b表示星期二 ... 根据用户的输入显示相依的信息。要求使用switch语句完成
package main

import "fmt"

func main() {

	//分析思路
	//1.定义一个变量接受字符
	//2.使用switch完成
	var key byte
	fmt.Println("请输入一个字符(a,b,c,d,e,f,g)：")
	fmt.Scanf("%c", &key)

	switch key {
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	case 'd':
		fmt.Println("周四")
	case 'e':
		fmt.Println("周五")
	case 'f':
		fmt.Println("周六")
	case 'g':
		fmt.Println("周日")
	default:
		fmt.Println("输入有误")
	}

	//switch 后可以不带表达式，类似 if--else分支来使用。
	var age int = 10

	switch {
	case age == 10:
		fmt.Println("age == 10")
	case age == 20:
		fmt.Println("age == 20")
	default:
		fmt.Println("age没有匹配到")
	}

	//case 中也可以对范围进行判断
	var score int = 98
	switch {
	case score > 90:
		fmt.Println("成绩优秀...")
	case score >= 70 && score <= 90:
		fmt.Println("成绩优良...")
	case score >= 60 && score < 70:
		fmt.Println("成绩及格...")
	default:
		fmt.Println("成绩不及格...")
	}

	//switch后也可以直接声明/定义一个变量，分号结束，不推荐

	switch grade := 90; {
	case grade > 90:
		fmt.Println("成绩优秀...")
	case grade >= 70 && grade <= 90:
		fmt.Println("成绩优良...")
	case grade >= 60 && grade < 70:
		fmt.Println("成绩及格...")
	default:
		fmt.Println("成绩不及格...")
	}

	//switch 的穿透 fallthrough
	var num int = 10
	switch num {
	case 10:
		fmt.Println("OK1")
		fallthrough
	case 20:
		fmt.Println("OK2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到....")
	}
	fmt.Println("main...over...")
}
