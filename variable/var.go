/**
 * 变量声明一般情况情况下使用var
 * var identifier type;
 * 常量使用const
 */

package main

import (
	"fmt"
	"unsafe"
)

func main2() {
	var a1 string = "字符串"
	fmt.Printf(a1)
	// 多个变量声明
	var b, c int = 1, 2
	fmt.Println(b, c)
	// 声明变量未初始化
	var d int  // 0
	var e bool // false
	fmt.Print(d, e)
	// 多个变量同时声明，可以为不同类型
	var m, n = 1, "字符串"
	const m1, n1 = 2, "字符串2"
	fmt.Println(m, n, m1, n1)
}

func main3() {
	// const 可作为枚举
	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	println(a, b, c)
}

func main() {
	// iota 特殊常量，可以被编辑器修改的常量
	// 第一个a 相当于0，后续自动累加。
	const (
		a = iota
		b = iota
		c //可以简写
	)
	println(a, b, c)
}
