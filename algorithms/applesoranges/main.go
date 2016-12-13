//https://www.hackerrank.com/challenges/apple-and-orange

package main

import "fmt"

func main() {
	var s, t, a, b, m, n int
	fmt.Scan(&s, &t, &a, &b, &m, &n)
	ad := make([]int, m)
	od := make([]int, n)
	oranges := 0
	apples := 0

	for i := 0; i < m; i++ {
		fmt.Scan(&ad[i])
		if a+ad[i] >= s && a+ad[i] <= t {
			apples++
		}
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&od[i])
		if b+od[i] >= s && b+od[i] <= t {
			oranges++
		}
	}
	fmt.Println(apples)
	fmt.Println(oranges)

}
