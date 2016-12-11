//https://www.hackerrank.com/challenges/circular-array-rotation

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n, q, k int
	fmt.Scan(&n, &q, &k)
	m := make([]int, k)
	var a []string

	var input string
	//use bufio
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()
		fmt.Println(len(input))
		fmt.Println(len(a))
		a = append(a, strings.Split(input, " ")...)
	}
	os.Exit(3)

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
	//a = append(a[n-1:n], a[0:n-1]...)
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
