//https://www.hackerrank.com/challenges/a-very-big-sum

package main

import (
	"fmt"
	"math/big"
)

func main() {

	var c int
	var j big.Int
	var sum big.Int
	fmt.Scan(&c)
	for i := 0; i < c; i++ {
		fmt.Scan(&j)
		sum.Add(&sum, &j)

	}
	fmt.Println(&sum)
}
