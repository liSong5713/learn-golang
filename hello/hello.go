package main // 可独立执行的程序
import (
	"fmt"
	"runtime"
) //引入fmt包

func main() {
	fmt.Println("hellow world \n", runtime.Version())
}
