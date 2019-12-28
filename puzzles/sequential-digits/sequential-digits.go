package main

import (
	"fmt"
	"math"
)

func leftmostDigit(n int) int {
	for n >= 10 {
		n = n / 10
	}

	return n
}

func numDigits(n int) int {
	var count int = 0
	for n > 0 {
		n = n / 10
		count = count + 1
	}
	return count
}

// assumes leftDigit and E are valid constraints for producing a sequential number, per validSequentialNumberConstraints
func getSequentialNumber(leftDigit int, E int) int {
	if leftDigit > 9 || E < 0 {
		return 0
	}

	var n int = leftDigit * int(math.Pow(10, float64(E)))
	return n + getSequentialNumber(leftDigit+1, E-1)
}

func validSequentialNumberConstraints(leftDigit int, E int) bool {
	return leftDigit+E <= 9
}

// https://leetcode.com/problems/sequential-digits/
// An integer has sequential digits if and only if each digit in the number is one more than the previous digit.
// Return a sorted list of all the integers in the range [low, high] inclusive that have sequential digits.
func sequentialDigits(low int, high int) []int {
	var startingDigit int = leftmostDigit(low)
	var startingE int = numDigits(low) - 1 // e.g. if low is in range 10 - 99, E is 1
	var maxE int = numDigits(high) - 1

	var result []int = make([]int, 0)

	E := startingE
	leftDigit := startingDigit
	for E <= maxE {
		for leftDigit < 10 {

			if validSequentialNumberConstraints(leftDigit, E) {
				var sequentialNumber int = getSequentialNumber(leftDigit, E)

				if low <= sequentialNumber && sequentialNumber <= high {
					result = append(result, sequentialNumber)
				}

			}

			leftDigit = leftDigit + 1
		}

		E = E + 1
		leftDigit = 1
	}

	return result
}

func main() {
	fmt.Println("hello")

	fmt.Println(leftmostDigit(123))
	fmt.Println(leftmostDigit(456))
	fmt.Println(numDigits(12345))
	fmt.Println(numDigits(10))
	fmt.Println(numDigits(11))

	fmt.Println(getSequentialNumber(1, 2)) // expects 123
	fmt.Println(getSequentialNumber(2, 4)) // expects 23456
	fmt.Println(getSequentialNumber(7, 2)) // expects 789

	fmt.Println(validSequentialNumberConstraints(7, 2)) // expects true
	fmt.Println(validSequentialNumberConstraints(8, 2)) // expects false

	fmt.Println(sequentialDigits(100, 300))    // expects [123, 234]
	fmt.Println(sequentialDigits(1000, 13000)) // expects [1234,2345,3456,4567,5678,6789,12345]
}
