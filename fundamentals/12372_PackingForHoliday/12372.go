package main

import "fmt"

func main() {
	var t, l, w, h int
	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		fmt.Scan(&l, &w, &h)
		if (l > 20) || (w > 20) || (h > 20) {
			fmt.Printf("Case %v: bad\n", i)
		} else {
			fmt.Printf("Case %v: good\n", i)
		}
	}
}
