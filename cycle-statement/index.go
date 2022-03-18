package main

import "fmt"

func cycel() {
	var a int = 1
	for a < 10 {
		fmt.Printf("a的值为:%d\n", a)
		if a == 5 {
			break
		}
		a++
	}
}

func main() {
	cycel()
}
