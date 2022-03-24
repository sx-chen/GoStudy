package main

import (
	"fmt"
)

func main() {
	// 1.打印九九乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			num := i * j
			fmt.Printf("%dx%d=%d\t", j, i, num)
		}
		// 2.打印换行符
		fmt.Println()
	}
}
