package main

import "fmt"

func main() {
	var a = 20
	var b = 10
	var area = a * b
	fmt.Println(area)
	// 关系运算符
	fmt.Println(a == b)
	fmt.Println(a > b)
	fmt.Println(a != b)
	// 条件
	if a != b {
		println("a不等于b")
	}
	// 逻辑运算符
	var c bool = false
	var d bool = true
	// 逻辑运算符只能接受bool类型，不能隐式转换
	var e = c && d
	println(e)
	if c || d {
		//  do
	}
	if !c {
	}
}
