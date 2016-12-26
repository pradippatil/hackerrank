//https://www.hackerrank.com/contests/w27/challenges/how-many-substrings

package main

import (
	"fmt"
	"index/suffixarray"
)

func main() {
	index := suffixarray.New([]byte("banana"))
	for i, v := range index.Bytes() {
		fmt.Println(i, &v)
	}
}
