package main

/* https://leetcode.com/problems/count-number-of-nice-subarrays/
Given an array of integers nums and an integer k. A subarray is called nice if there are k odd numbers on it.

Return the number of nice sub-arrays.



Example 1:

Input: nums = [1,1,2,1,1], k = 3
Output: 2
Explanation: The only sub-arrays with 3 odd numbers are [1,1,2,1] and [1,2,1,1].

*/

import "fmt"

// i is the index on indices. padding left is the space between
// the nums index at indices[i] and either the beginning of the nums array, or the next
// odd number to the left
//
// the reason paddingLeft matters is that if you find a minimum subarray with the correct number of odd numbers.
// paddingLeft represents the number of *additional* subarrays with the correct number of odd numbers
// you can derive by simply expanding the subarray to the left. e.g. with the array
// [2, 2, 2, 1, 1, 1], k = 3
// the minimum valid subarrayy is [1, 1, 1] -- but you can add the twos to the left
// without changing the number of odds. e.g. [2, 1, 1, 1] and [2, 2, 1, 1, 1] are also valid.
func paddingLeft(indices []int, i int) int {

	// case where this is the first index, so rather than comparing distance to next odd to the left
	// (which DNE), compare to start of nums array, equivalent to the index itself indices[i]
	if i == 0 {
		return indices[i]
	}

	// distance between this odd number and the previous odd number
	return indices[i] - (1 + indices[i-1])
}

func isOdd(n int) bool {
	return n%2 == 1
}

func numberOfSubarrays(nums []int, k int) int {
	oddIndices := make([]int, 0)
	var leftmostOdd int = 0
	var count int = 0
	var oddCount int = 0

	for i, n := range nums {
		if isOdd(n) {
			oddIndices = append(oddIndices, i)
			oddCount++

			// here we effectively remove the leftmost oddnumber once we exceed the desired number of odds
			if oddCount > k {
				leftmostOdd++
				oddCount--
			}

		}

		if oddCount == k {
			// the 1 represents the minum subarray we just found. and paddingLeft represents the
			// trivial expansions of that subarray -- if any -- which we can make by simply
			// adding numbers to the left of the minimum subarray until we hit an odd number
			count = count + 1 + paddingLeft(oddIndices, leftmostOdd)
		}
	}

	return count
}

func main() {
	nums1 := []int{1, 1, 2, 1, 1}
	k1 := 3
	fmt.Println(numberOfSubarrays(nums1, k1)) // expect 2

	nums2 := []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}
	k2 := 2
	fmt.Println(numberOfSubarrays(nums2, k2)) // expect 16

}
