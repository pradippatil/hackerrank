//https://www.hackerrank.com/challenges/staircase

package main

import "fmt"
import "strings"

func main() {
	var count int
	fmt.Scan(&count)
	for i := 1; i <= count; i++ {
		fmt.Printf("%[1]*[2]s\n", count, strings.Repeat("#", i))
	}
}
