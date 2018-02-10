//To complete this challenge, you must save a line of input from stdin to a variable, print Hello, World. on a single line, and finally print the value of your variable on a second line.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var inputString string
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); ok {
		inputString = scanner.Text()
	}
	fmt.Printf("%s\n%s", "Hello, World.", inputString)
}
