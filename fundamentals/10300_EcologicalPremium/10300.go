package main

import (
	"fmt"
	"math"
)

func main() {
	var n, f, prem int
	var a, b, c float64
	fmt.Scan(&n)
	if (n >= 20) || (n <= 0) {
		fmt.Println("Invalid number of test cases")
		fmt.Scan(&n)
	}
	for n > 0 {
		fmt.Scan(&f)
		if (f <= 0) || (f >= 20) {
			fmt.Println("Invalid number of farmers in test case")
			fmt.Scan(&f)
		}
		prem = 0
		for f > 0 {
			fmt.Scan(&a, &b, &c)
			prem += int(math.Round(((a / b) * c) * b))
			f--
		}
		fmt.Println(prem)
		n--
	}
}
