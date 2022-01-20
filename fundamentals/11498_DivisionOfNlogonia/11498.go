package main

import "fmt"

func main() {
	var t, k, m int
	fmt.Scan(&t)
	fmt.Scan(&k, &m)
	for t != 0 {
		for i := 0; i < t; i++ {
			if t == 0 {
				break
			}
			var x, y int
			fmt.Scan(&x, &y)
			if (x > k) && (y > m) {
				fmt.Println("NE")
			} else if (x > k) && (y < m) {
				fmt.Println("SE")
			} else if (x < k) && (y > m) {
				fmt.Println("NO")
			} else if (x < k) && (y < m) {
				fmt.Println("SO")
			} else if (x == k) || (y == m) {
				fmt.Println("divisa")
			}
		}
		fmt.Scan(&t)
		fmt.Scan(&k, &m)
	}
}
