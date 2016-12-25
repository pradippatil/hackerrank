//https://www.hackerrank.com/challenges/nim-game-1

package main

import "fmt"

func grundy(n int64) int64 {
	if n == 0 {
		return 0
	}
	return n
}

func main() {
	var g, n int
	var nimSum int64
	fmt.Scan(&g)
	var p = make([][]int64, g)

	for i := 0; i < g; i++ {
		fmt.Scan(&n)
		p[i] = make([]int64, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&p[i][j])
		}
	}

	for i := 0; i < g; i++ {
		n := len(p[i])
		G := make([]int64, n)
		nimSum = 0
		for j := 0; j < n; j++ {
			G[j] = grundy(p[i][j])
		}

		//XOR of Grundy numbers
		for _, v := range G {
			nimSum ^= v
		}

		if nimSum != 0 {
			fmt.Println("First")
		} else {
			fmt.Println("Second")
		}

	}
}
