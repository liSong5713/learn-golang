/*
* 格式 func func_name([params list]) (return_types){
*
*}
**/
package main

import "fmt"

// 1、函数返回多个值
func twoReturnType() (int, string) {
	var a, b = 10, "v_v"
	return a, b
}

//  2、go函数的参数为值传递
var x int = 20
var y = "oly"

func paramCheck(x int, y string) int {
	x = 30
	return x
}

// 3、go 函数可以作为另外一个函数的实参
var f1 = func() int {
	return 1
}

func f2() {
	//  f1作为Println的参数
	fmt.Println(f1())
}

// 4、 函数作为参数传递，实现回调

type cb func(int) int

var f3 cb = func(x int) int {
	return x
}

func testCallBack(x int, f cb) int {
	return f(x)
}

// 5、函数闭包
func sequence() func() int {
	var a = 0
	return func() int {
		a += 1
		return a
	}
}

var nextNumber = sequence()

func main() {
	var v = paramCheck(x, y)
	fmt.Println("origin x:", x, "result:", v)

	fmt.Println("callback result:", testCallBack(x, f3))
	fmt.Println("nextNumber ++1:", nextNumber())
	fmt.Println("nextNumber ++2:", nextNumber())
}
