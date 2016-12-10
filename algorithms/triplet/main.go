//https://www.hackerrank.com/challenges/compare-the-triplets

package main

import "fmt"

func main() {
	a := make([]int, 3)
	b := make([]int, 3)
	var sa, sb int
	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < len(b); i++ {
		fmt.Scan(&b[i])
	}

	for i := range a {
		if a[i] > b[i] {
			sa++
		} else if b[i] > a[i] {
			sb++
		}
	}
	fmt.Println(sa, sb)

}
