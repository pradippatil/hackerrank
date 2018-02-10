//https://www.hackerrank.com/challenges/30-2d-array

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string
	m := make([][]int, 6)
	r := bufio.NewReader(os.Stdin)
	for i := 0; i < len(m); i++ {
		input, _ = r.ReadString('\n')
		m[i] = make([]int, 6)
		fmt.Sscanf(input, "%d %d %d %d %d %d\n", &m[i][0], &m[i][1], &m[i][2], &m[i][3], &m[i][4], &m[i][5])

	}
	fmt.Println(m)
	var maxValue int
	var sum int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum = m[i][j] + m[i][j+1] + m[i][j+2] + m[i+1][j+1] + m[i+2][j] + m[i+2][j+1] + m[i+2][j+2]

			//reset maxValue at first sum or when sum exceeds existing maxValue
			if (i == 0 && j == 0) || sum > maxValue {
				maxValue = sum
			}

		}
	}
	fmt.Println(maxValue)
}
