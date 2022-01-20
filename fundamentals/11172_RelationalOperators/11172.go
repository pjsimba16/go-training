package main

import "fmt"

func main() {
	var t, a, b int
	fmt.Scan(&t)
	for t > 0 {
		fmt.Scan(&a, &b)
		if a > b {
			fmt.Println(">")
		} else if b > a {
			fmt.Println("<")
		} else {
			fmt.Println("=")
		}
		t--
	}
}
