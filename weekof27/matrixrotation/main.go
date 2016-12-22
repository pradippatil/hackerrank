//https://www.hackerrank.com/contests/w27/challenges/hackonacci-matrix-rotations
package main

import "fmt"

//M is multiplication matrix
var M = [][]int64{
	{1, 2, 3},
	{1, 0, 0},
	{0, 1, 0},
}

var N = [][]int64{
	{3, 0, 0},
	{2, 0, 0},
	{1, 0, 0},
}

func multiply(a, b [][]int64) [][]int64 {
	n := len(a)
	mul := make([][]int64, n)
	for i := 0; i < n; i++ {
		mul[i] = make([]int64, n)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				mul[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	//fmt.Println("Multiplication of", a, " and ", b, "is", mul)
	return mul
}

//power returns matrix^n
func power(a [][]int64, n int64) [][]int64 {
	if n == 1 {
		return a
	}

	a = power(a, n/2)
	a = multiply(a, a)

	if n%2 != 0 {
		a = multiply(a, M)
	}
	//fmt.Println("Power of", a, " to ", n, "is", a)

	return a
}

func hackonacci(n int64) int64 {
	if n == 1 || n == 2 || n == 3 {
		return n
	}
	return multiply(power(M, n-3), N)[0][0]
}

func rotate(m [][]int) [][]int {
	n := len(m)
	r := make([][]int, n)
	for i := 0; i < n; i++ {
		r[i] = make([]int, n)
		for j := 0; j < n; j++ {
			r[i][j] = m[n-j-1][i]
		}
	}
	return r
}

func main() {
	var n, q, i, j int64
	var h = make(map[int64]int64)
	fmt.Scan(&n, &q)
	var angles = make([]int, q)
	for i = 0; i < q; i++ {
		fmt.Scan(&angles[i])
	}
	var m = make([][]int, n)

	if (n+1)%7 == 0 {
		fmt.Println(0)
		return
	}

	for i = 1; i <= n; i++ {
		m[i-1] = make([]int, n)
		for j = 1; j <= n; j++ {
			if i > 7 || j > 7 {
				m[i-1][j-1] = m[(i-1)%7][(j-1)%7]
			} else {
				index := (i * j) * (i * j)
				//fmt.Println("index is", index)
				if _, ok := h[index]; !ok {
					h[index] = hackonacci(index)
				}
				if h[index]%2 == 0 {
					m[i-1][j-1] = 1
				}
			}

		}
	}
	//fmt.Println(m)
	var count = make(map[int]int64)
	for _, v := range angles {
		v = v % 360
		if _, ok := count[v]; !ok {
			if v == 0 || v == 360 {
				count[v] = 0
			} else {

				r := make([][]int, n)
				copy(r, m)
				for i := 0; i < v/90; i++ {
					r = rotate(r)
				}
				for i = 0; i < n; i++ {
					for j = 0; j < n; j++ {
						if m[i][j] != r[i][j] {
							count[v]++
						}

					}

				}
			}

		}
		fmt.Println(count[v])
	}
}
