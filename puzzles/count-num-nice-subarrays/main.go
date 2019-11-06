package main

/* https://leetcode.com/problems/count-number-of-nice-subarrays/
Given an array of integers nums and an integer k. A subarray is called nice if there are k odd numbers on it.

Return the number of nice sub-arrays.



Example 1:

Input: nums = [1,1,2,1,1], k = 3
Output: 2
Explanation: The only sub-arrays with 3 odd numbers are [1,1,2,1] and [1,2,1,1].

*/

// i is the index on indices. padding left is the space between
// the nums index at indices[i] and either the beginning of the nums array, or the next
// odd number to the left
func paddingLeft(indices []int, i int) int {

	// case where this is the first index, so rather than comparing distance to next odd to the left
	// (which DNE), compare to start of nums array, equivalent to the index itself indices[i]
	if i == 0 {
		return indices[i]
	}

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

	for n, i := range nums {
		if isOdd(n) {
			oddIndices = append(oddIndices, i)
			oddCount++

			if oddCount > k {
				leftmostOdd++
				oddCount--
			}

		}

		if oddCount == k {
			count = count + 1 + paddingLeft(oddIndices, leftmostOdd)
		}
	}

	return count
}

func main() {

}
