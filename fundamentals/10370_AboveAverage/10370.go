package main

import "fmt"

func main() {
	var c, n int
	fmt.Scan(&c)
	for c > 0 {
		var grades_ls []int
		var avg, sum, above_avg_perc, above_avg_count float64
		above_avg_count = 0
		fmt.Scan(&n)
		for n > 0 {
			var grade int
			fmt.Scan(&grade)
			grades_ls = append(grades_ls, grade)
			n--
		}
		for i := 0; i < len(grades_ls); i++ {
			sum += float64(grades_ls[i])
		}
		avg = sum / float64(len(grades_ls))
		for j := 0; j < len(grades_ls); j++ {
			if float64(grades_ls[j]) > avg {
				above_avg_count += 1
			}
		}
		above_avg_perc = (above_avg_count / float64(len(grades_ls))) * 100
		fmt.Printf("%3.3f%%\n", above_avg_perc)
		c--
	}
}
