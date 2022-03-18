package main

import (
	"fmt"
	"strconv"
)

func int2String() {
	var a int = 10
	// int=>string
	var as = strconv.Itoa(a)
	fmt.Println(as)
	// int64 => string
	var i64 int64 = 12
	var i64s = strconv.FormatInt(i64, 2) // 参数二为进制
	fmt.Println(i64s)
}

func string2Int() {
	var s, err = strconv.Atoi("124")
	fmt.Println(s, err)
	// m2 有符号整型
	var s2, err2 = strconv.ParseInt("123", 10, 64)
	fmt.Println(s2, err2)

	// m3 无符号整型
	var s3, err3 = strconv.ParseUint("125", 10, 64)
	fmt.Println(s3, err3)
}

func main() {
	string2Int()

}
