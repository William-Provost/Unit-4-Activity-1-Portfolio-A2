// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-24
// Fileoverview: This program asks the user for two odd numbers and prints all even numbers between them.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// variable declaration
	var startString string
	var endString string
	var startValue int
	var endValue int

	reader := bufio.NewReader(os.Stdin)

	// ask for starting value
	fmt.Print("Please enter a starting value (it must be an odd number): ")
	startString, _ = reader.ReadString('\n')
	startString = strings.TrimSpace(startString)
	startValue, _ = strconv.Atoi(startString)

	// ask for ending value
	fmt.Print("Please enter an ending value (it must be an odd number): ")
	endString, _ = reader.ReadString('\n')
	endString = strings.TrimSpace(endString)
	endValue, _ = strconv.Atoi(endString)

	fmt.Println("\nEven numbers between your values:")

	// loop between values
	for i := startValue; i <= endValue; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", i)
		}
	}

	fmt.Println("\nDone.")
}
