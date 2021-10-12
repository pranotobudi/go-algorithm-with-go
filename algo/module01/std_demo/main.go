package main

import (
	"fmt"
	"algo/module01"
)

func main() {
	var n int
	fmt.Scanf("%d %f\n", &n) // we need to put %f\n because Scanf rejects \r\n at end of line on Windows
	// fmt.Println("total: ", n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scanf("%d %d %f\n", &a, &b) // we need to put %f\n because Scanf rejects \r\n at end of line on Windows
		fmt.Println("a b", a, b)
		gcd := module01.GCD(a, b)
		fmt.Println(gcd)
	}
}
