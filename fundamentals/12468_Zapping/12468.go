package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for (a != -1) && (b != -1) {
		var diff1, diff2 int
		if a >= b {
			diff1 = a - b
			diff2 = (100 + b) - a
		} else {
			diff1 = b - a
			diff2 = (100 + a) - b
		}
		if diff1 > diff2 {
			fmt.Println(diff2)
		} else {
			fmt.Println(diff1)
		}
		fmt.Scan(&a, &b)
	}
}
