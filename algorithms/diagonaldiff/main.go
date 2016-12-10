//https://www.hackerrank.com/challenges/diagonal-difference

package main

import "fmt"

func main() {
	var count, d, d1, d2 int
	fmt.Scan(&count)
	a := make([][]int, count)
	for i := 0; i < count; i++ {
		a[i] = make([]int, count)
		for j := 0; j < count; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	for i := 0; i < count; i++ {
		d1 += a[i][i]
		d2 += a[i][count-i-1]
	}
	d = d1 - d2

	//Need absolute difference
	if d < 0 {
		d = -d
	}
	fmt.Println(d)

}
