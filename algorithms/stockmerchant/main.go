//https://www.hackerrank.com/challenges/sock-merchant

package main

import "fmt"

func main() {

	var n, c, pairs int
	var socks = make(map[int]int)
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&c)
		//Store count of each color
		socks[c]++
	}

	for _, v := range socks {
		//Count number of pairs of same color
		pairs += v / 2
	}
	fmt.Println(pairs)

}
