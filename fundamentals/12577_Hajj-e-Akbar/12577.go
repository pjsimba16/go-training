package main

import "fmt"

func main() {
	var h string
	var count int
	count = 1
	fmt.Scan(&h)
	for h != "*" {
		if h == "Hajj" {
			fmt.Printf("Case %v: Hajj-e-Akbar\n", count)
		} else if h == "Umrah" {
			fmt.Printf("Case %v: Hajj-e-Asghar\n", count)
		} else {
			fmt.Printf("Invalid Input\n")
		}
		fmt.Scan(&h)
		count++
	}
}
