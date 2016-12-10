//https://www.hackerrank.com/challenges/plus-minus

package main

import "fmt"

func main() {
	var count, p, n, z int
	fmt.Scan(&count)
	a := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&a[i])
	}
	for _, v := range a {
		switch {
		case v < 0:
			n++
		case v > 0:
			p++
		case v == 0:
			z++
		}
	}
	fmt.Printf("%.6f\n%.6f\n%.6f", float32(p)/float32(count), float32(n)/float32(count), float32(z)/float32(count))

}
