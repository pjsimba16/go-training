package main

import "fmt"

func main() {
	var T, a, b, c, counter int
	fmt.Scan(&T)
	counter = 1
	for T > 0 {
		fmt.Scan(&a, &b, &c)
		if ((a < b) && (a > c)) || ((a > b) && (a < c)) {
			fmt.Printf("Case %v: %v\n", counter, a)
		} else if ((b < a) && (b > c)) || ((b > a) && (b < c)) {
			fmt.Printf("Case %v: %v\n", counter, b)
		} else {
			fmt.Printf("Case %v: %v\n", counter, c)
		}
		T--
		counter++
	}
}
