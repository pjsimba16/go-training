package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var num string
		fmt.Scan(&num)
		if len(num) == 5 {
			fmt.Println(3)
		} else if len(num) == 3 {
			if (string(num[0]) == "o" && string(num[1]) == "n") ||
				(string(num[0]) == "o" && string(num[2]) == "e") ||
				(string(num[1]) == "n" && string(num[2]) == "e") {
				fmt.Println(1)
			} else if (string(num[0]) == "t" && string(num[1]) == "w") ||
				(string(num[0]) == "t" && string(num[2]) == "o") ||
				(string(num[1]) == "w" && string(num[2]) == "o") {
				fmt.Println(2)
			}
		} else {
			fmt.Println("invalid input")
		}
	}
}
