//https://www.hackerrank.com/challenges/30-binary-numbers

package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)

	var remainder, max, counter int

	for input > 0 {
		remainder = input % 2
		input = input / 2
		if remainder == 1 {
			counter++
			if counter >= max {
				max = counter
			}
		} else {
			counter = 0

		}
	}

	fmt.Println(max)
}
