//https://www.hackerrank.com/challenges/30-dictionaries-and-maps

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	var N int

	phoneBook := make(map[string]string)

	fmt.Scan(&N)
	if N < 1 && N >= int(math.Pow(10, 5)) {
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < N; i++ {
		if ok := scanner.Scan(); ok {
			input := strings.Split(scanner.Text(), " ")
			phoneBook[input[0]] = input[1]
		}

	}
	var inputNames []string
	for scanner.Scan() {
		inputNames = append(inputNames, scanner.Text())
	}

	if len(inputNames) < 1 && len(inputNames) >= int(math.Pow(10, 5)) {
		return
	}
	for index := range inputNames {
		val, ok := phoneBook[inputNames[index]]
		if ok {
			fmt.Printf("%s=%s\n", inputNames[index], val)
		} else {
			fmt.Println("Not found")
		}
	}

}
