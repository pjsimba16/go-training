package main

import (
	"fmt"
	"sort"
)

func main() {
	var x int
	var x_ls []int
	_, err := fmt.Scan(&x)
	for err == nil {
		x_ls = append(x_ls, x)
		sort.Ints(x_ls)
		if len(x_ls)%2 == 0 {
			fmt.Println((x_ls[(len(x_ls)/2)-1] + x_ls[len(x_ls)/2]) / 2)
		} else {
			fmt.Println(x_ls[len(x_ls)/2])
		}
		_, err = fmt.Scan(&x)
	}
}
