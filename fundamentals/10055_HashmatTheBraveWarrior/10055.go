package main

import "fmt"

func main() {
	var h, o int
	_, err := fmt.Scan(&h, &o)
	for err == nil {
		if o > h {
			fmt.Println(o - h)
		} else {
			fmt.Println(h - o)
		}
		_, err = fmt.Scan(&h, &o)
	}
}
