package main

import "fmt"

func main() {
	var s string
	var count int
	count = 1
	fmt.Scan(&s)
	for s != "#" {
		if s == "HELLO" {
			fmt.Printf("Case %v: ENGLISH\n", count)
		} else if s == "HOLA" {
			fmt.Printf("Case %v: SPANISH\n", count)
		} else if s == "HALLO" {
			fmt.Printf("Case %v: GERMAN\n", count)
		} else if s == "BONJOUR" {
			fmt.Printf("Case %v: FRENCH\n", count)
		} else if s == "CIAO" {
			fmt.Printf("Case %v: ITALIAN\n", count)
		} else if s == "ZDRAVSTVUJTE" {
			fmt.Printf("Case %v: RUSSIAN\n", count)
		} else {
			fmt.Printf("Case %v: UNKNOWN\n", count)
		}
		fmt.Scan(&s)
		count++
	}
}
