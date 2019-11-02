package main

import "fmt"

// puzzle from https://leetcode.com/problems/split-a-string-in-balanced-strings/
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

		// idea of the algorithm is that if we reach a point where we have seen an equal
		// number of Ls and Rs since the last split, then we have reached the earliest possible valid split
		// for this part of the string
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
