package main

import "fmt"

func main() {
	var T, n, max int
	fmt.Scan(&T)
	for i := 1; i <= T; i++ {
		var s_ls []int
		fmt.Scan(&n)
		for j := 1; j <= n; j++ {
			var s int
			fmt.Scan(&s)
			s_ls = append(s_ls, s)
		}
		for k := 1; k < len(s_ls); k++ {
			if s_ls[0] < s_ls[k] {
				s_ls[0] = s_ls[k]
			}
		}
		max = s_ls[0]
		fmt.Printf("Case %v: %v\n", i, max)
	}
}
