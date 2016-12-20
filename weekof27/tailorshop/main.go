//https://www.hackerrank.com/contests/w27/challenges/tailor-shop

package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, p, total int
	c := make(map[int]bool)
	counts := make(map[int]int)
	fmt.Scan(&n, &p)
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&prices[i])
	}
	sort.Ints(prices)

	for i := 0; i < len(prices); i++ {
		count := 0
		c[count] = true
		if prices[i]%p == 0 {
			count += prices[i] / p
		} else {
			count += prices[i]/p + 1

		}
		if i > 0 && c[count] == true {
			count = counts[i-1] + 1
		}
		c[count] = true
		counts[i] = count
		total += count
	}

	fmt.Println(total)
}
