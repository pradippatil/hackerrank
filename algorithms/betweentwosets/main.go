/*
https://www.hackerrank.com/challenges/between-two-sets
O(n log(n)) solution.
1. find the LCM of all the integers of array A.
2. find the GCD of all the integers of array B.
3. Count the number of multiples of LCM that evenly divides the GCD.
*/
package main

import "fmt"

//gcd returns greatest common divisor (GCD) for given 2 numbers
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

//lcm returns least common multiple (LCM)  for given 2 numbers
func lcm(a, b int) int {
	return a / gcd(a, b) * b

}

func main() {
	var n, m, l, g, count int
	fmt.Scan(&n, &m)
	A := make([]int, n)
	B := make([]int, m)

	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&B[i])
	}

	//Find LCM for set A
	l = A[0]
	for i := 1; i < n; i++ {
		l = lcm(l, A[i])
	}

	//Find GCD for set B
	g = B[0]
	for i := 0; i < m; i++ {
		g = gcd(g, B[i])

	}

	//Find multiples of LCM which can divide GCM
	c := 1
	for l*c <= g {
		if g%(l*c) == 0 {
			count++
		}
		c++
	}
	fmt.Println(count)

}
