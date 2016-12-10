//https://www.hackerrank.com/challenges/simple-array-sum

package main

import "fmt"

func main() {
	var size, sum int
	fmt.Scan(&size)
	a := make([]int, size)
	for i := 0; i < size; i++ {
		//We could sum all numbers here, but since this tutorial is to use array, let's store given values in array
		fmt.Scan(&a[i])
	}
	for i := range a {
		sum += a[i]
	}
	fmt.Println(sum)

}
