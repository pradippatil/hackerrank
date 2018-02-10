//https://www.hackerrank.com/challenges/30-loops

package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", N, i, N*i)
	}

}
