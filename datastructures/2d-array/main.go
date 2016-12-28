//https://www.hackerrank.com/challenges/2d-array

package main

import "fmt"

func main() {
	n := 6
	a := make([][]int, n)

	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	// calculate hourglass sum
	max := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := a[i][j] + a[i][j+1] + a[i][j+2] + a[i+1][j+1] + a[i+2][j] + a[i+2][j+1] + a[i+2][j+2]

			if (i == 0 && j == 0) || sum > max {
				max = sum
			}
		}
	}
	fmt.Println(max)
}
