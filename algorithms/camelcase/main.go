//https://www.hackerrank.com/challenges/camelcase

package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)
	var wc int

	//count word starting with lower letter
	if len(r) > 0 {
		wc++
	}
	for i := range r {
		if unicode.IsUpper(r[i]) == true {
			wc++
		}
	}
	fmt.Println(wc)

}
