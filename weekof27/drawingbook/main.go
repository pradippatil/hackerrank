//https://www.hackerrank.com/contests/w27/challenges/drawing-book

package main

import (
	"fmt"
)

func main() {
	var n, p int
	fmt.Scan(&n, &p)

	if p/2 < (n-p)/2 {
		// If page number is in first half, it would need to turn half the pages from start
		fmt.Println(p / 2)
	} else {
		// If page number is in second half, it would need to turn half the pages from end
		fmt.Println((n - p) / 2)
	}

}
