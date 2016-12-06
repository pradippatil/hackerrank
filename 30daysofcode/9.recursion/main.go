//https://www.hackerrank.com/challenges/30-recursion

package main

import "fmt"

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

func main() {

	var in int
	//input from terminal
	fmt.Scan(&in)
	fmt.Println(factorial(in))

}
