package main

import "fmt"

func main() {
	var stockcode = 123
	var endate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	url = "1212"
	// url=1  //违法操作 类型不一致
	shortV := 1 //快捷声明
	var target_url = fmt.Sprintf(url, stockcode, endate)
	fmt.Print(target_url)
}
