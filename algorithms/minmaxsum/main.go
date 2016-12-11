//https://www.hackerrank.com/challenges/mini-max-sum

package main

import "fmt"
import "sort"

func main() {
	a := make([]int, 5)
	var minSum, maxSum int64
	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	for _, v := range a[:4] {
		minSum += int64(v)
	}
	for _, v := range a[1:5] {
		maxSum += int64(v)
	}
	fmt.Println(minSum, maxSum)
}
