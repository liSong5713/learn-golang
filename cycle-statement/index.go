package main

import "fmt"

func cycel() {
	var a int = 1
	// way1
	for a < 10 {
		fmt.Printf("a的值为:%d\n", a)
		if a == 5 {
			break
		}
		a++
	}
	// way2
	for b := 0; b < 10; b++ {
	}
}

func main() {
	cycel()
}
