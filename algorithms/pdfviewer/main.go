//https://www.hackerrank.com/challenges/designer-pdf-viewer
package main

import "fmt"

func main() {
	a := "abcdefghijklmnopqrstuvwxyz"
	ah := make([]int, 26)
	am := make(map[byte]int, 26)
	var in string
	//store height in a map
	for i := 0; i < len(a); i++ {
		fmt.Scan(&ah[i])
		am[a[i]] = ah[i]
	}
	fmt.Scan(&in)

	//calculate max height in given string
	max := 0
	for i := 0; i < len(in); i++ {
		if am[in[i]] > max {
			max = am[in[i]]
		}
	}
	fmt.Println(1 * len(in) * max)

}
