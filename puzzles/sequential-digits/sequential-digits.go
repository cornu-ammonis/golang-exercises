package main

import (
	"fmt"
)

func leftmostDigit(n int) int {
	for n >= 10 {
		n = n / 10
	}

	return n
}

// https://leetcode.com/problems/sequential-digits/
// An integer has sequential digits if and only if each digit in the number is one more than the previous digit.
// Return a sorted list of all the integers in the range [low, high] inclusive that have sequential digits.
func sequentialDigits(low int, high int) []int {
	return []int{}
}

func main() {
	fmt.Println("hello")

	fmt.Println(leftmostDigit(123))
	fmt.Println(leftmostDigit(456))
}