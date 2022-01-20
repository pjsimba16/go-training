package main

import "fmt"

func main() {
	var t, r1, r2, sum int
	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		sum = 0
		fmt.Scan(&r1, &r2)
		for j := r1; j <= r2; j++ {
			if j%2 == 1 {
				sum += j
			}
		}
		fmt.Printf("Case %v: %v\n", i, sum)
	}
}
