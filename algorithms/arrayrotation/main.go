//https://www.hackerrank.com/challenges/circular-array-rotation

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, q, k int
	fmt.Scan(&n, &q, &k)
	m := make([]int, k)
	var a []string

	//use bufio
	r := bufio.NewReader(os.Stdin)
	input, isPrefix, err := r.ReadLine()
	fmt.Println("Read", string(input))
	if isPrefix {
		fmt.Println("buffer size to small")
	}
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(len(a))
	//a = append(a, strings.Split(input, " ")...)

	//Read array
	//for i := 0; i < n; i++ {
	//fmt.Scan(&a[i])
	//}
	//Read indexes to find after q rotations
	for i := 0; i < k; i++ {
		fmt.Scan(&m[i])
	}

	fmt.Println(len(a))
	fmt.Println("Exiting")
	os.Exit(1)

	//Rotate array q times

	//for i := 0; i < q; i++ {
	//a = append(a[n-1:], a[0:n-1]...)
	//}

	//Find given indexes
	for _, v := range m {
		if v >= q {
			fmt.Println(a[v-q])
		} else if v < q {
			fmt.Println(a[n-(v+1)])
		}
	}

}
