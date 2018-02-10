//https://www.hackerrank.com/challenges/30-review-loop

package main

import "fmt"
import "bytes"

func main() {
	var T int
	var strings []string
	var oddString bytes.Buffer
	var evenString bytes.Buffer
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var s string
		fmt.Scan(&s)
		strings = append(strings, s)
	}

	for _, S := range strings {
		for i := 0; i < len(S); i++ {
			if (i+2)%2 == 0 {
				evenString.WriteString(string(S[i]))
			} else {
				oddString.WriteString(string(S[i]))
			}
		}
		fmt.Printf("%s %s\n", evenString.String(), oddString.String())
		evenString.Reset()
		oddString.Reset()
	}
}
