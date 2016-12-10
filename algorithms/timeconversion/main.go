//https://www.hackerrank.com/challenges/time-conversion
package main

import (
	"fmt"
	"time"
)

func main() {
	var s string
	fmt.Scan(&s)
	//Go uses different approach for parsing and formatting using specific layout string - see time package
	t, _ := time.Parse("03:04:05PM", s)
	fmt.Println(t.Format("15:04:05"))
}
