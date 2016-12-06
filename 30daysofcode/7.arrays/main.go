//https://www.hackerrank.com/challenges/30-arrays
package main

import "fmt"

func main() {
	var N int
	var A []int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var n int
		fmt.Scan(&n)
		A = append(A, n)
	}

	for i := len(A) - 1; i >= 0; i-- {
		fmt.Print(A[i], " ")
	}
}
