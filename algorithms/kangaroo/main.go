//https://www.hackerrank.com/challenges/kangaroo

package main

import "fmt"

func main() {
	var x1, x2, v1, v2 int
	fmt.Scan(&x1, &v1, &x2, &v2)
	cross := false

	if x1 == x2 && v1 == v2 {
		//if both start at same location and same speed = YES
		cross = true
	} else if x2 > x1 && v1 > v2 && (x2-x1)%(v1-v2) == 0 {
		// if first started early && faster && if starting distance difference is multiple of velocity diff = YES
		cross = true
	} else if x1 > x2 && v2 > v1 && (x1-x2)%(v2-v1) == 0 {
		// same as above for second kangaroo
		cross = true
	}

	if cross {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
