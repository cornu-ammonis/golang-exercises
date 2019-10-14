package main

import "fmt"

func balancedStringSplit(s string) int {
	var lCount int = 0
	var rCount int = 0
	var balancedCount int = 0

	for _, c := range s {
		if c == 'L' {
			lCount++
		} else {
			rCount++
		}

		if lCount == rCount {
			balancedCount++
			lCount = 0
			rCount = 0
		}
	}

	return balancedCount
}

func main() {

	fmt.Println(balancedStringSplit("RLRRLLRLRL")) // expects 4
	fmt.Println(balancedStringSplit("RLLLLRRRLR")) // expects 3
	fmt.Println(balancedStringSplit("LLLLRRRR"))   // expects 1
}
