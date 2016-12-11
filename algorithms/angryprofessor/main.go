//https://www.hackerrank.com/challenges/angry-professor

package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	cancelled := make([]string, t)

	for i := 0; i < t; i++ {
		var total, threshold, onTime int
		fmt.Scan(&total)
		fmt.Scan(&threshold)
		arrival := make([]int, total)
		for j := 0; j < total; j++ {
			fmt.Scan(&arrival[j])
			if arrival[j] <= 0 {
				onTime++
			}
		}

		if onTime >= threshold {
			cancelled[i] = "NO"
		} else {
			cancelled[i] = "YES"
		}
	}

	for i := 0; i < t; i++ {
		fmt.Println(cancelled[i])
	}
}
