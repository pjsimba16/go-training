package main

import "fmt"

func main() {
	var T, counter int
	var a, b, c int32
	counter = 1
	fmt.Scan(&T)
	for T > 0 {
		fmt.Scan(&a, &b, &c)
		if ((a + b) <= c) || ((a + c) <= b) || ((b + c) <= a) {
			fmt.Printf("Case %v: Invalid\n", counter)
		} else if (a == b) && (a == c) && (b == c) {
			fmt.Printf("Case %v: Equilateral\n", counter)
		} else if (a == b) || (a == c) || (b == c) {
			fmt.Printf("Case %v: Isosceles\n", counter)
		} else {
			fmt.Printf("Case %v: Scalene\n", counter)
		}
		T--
		counter++
	}
}
