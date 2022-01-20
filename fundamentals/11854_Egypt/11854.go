package main

import "fmt"

func main() {
	var a, b, c, a2, b2, c2 int
	var t bool
	t = true
	fmt.Scan(&a, &b, &c)
	for t == true {
		a2 = a * a
		b2 = b * b
		c2 = c * c
		if (a == 0) && (b == 0) && (c == 0) {
			break
		} else if ((a2 + b2) == c2) || ((c2 + b2) == a2) || ((a2 + c2) == b2) {
			fmt.Println("right")
		} else {
			fmt.Println("wrong")
		}
		fmt.Scan(&a, &b, &c)
	}
}
