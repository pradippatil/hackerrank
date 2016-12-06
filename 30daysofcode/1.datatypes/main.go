//https://www.hackerrank.com/challenges/30-data-types

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var i uint32 = 4
	var d float32 = 4.0
	var s string = "HackerRank "
	scanner := bufio.NewReader(os.Stdin)

	var i2 uint32
	var d2 float32
	var s2 string
	// Scan int and float
	fmt.Scan(&i2)
	fmt.Scan(&d2)

	//fmt.Scan doesn't work for space separated strings. Use bufio.NewReader
	s2, _ = scanner.ReadString('\n')
	fmt.Printf("%d\n%.1f\n%s%s\n", i+i2, d+d2, s, s2)

}
