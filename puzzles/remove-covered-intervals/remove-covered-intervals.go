package main

import (
	"fmt"
	"sort"
)

// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func remove(slice [][]int, s int) [][]int {
	return append(slice[:s], slice[s+1:]...)
}

// tests whether potential covers target
func covers(target []int, potential []int) bool {
	return potential[0] <= target[0] && target[1] <= potential[1]
}

// https://leetcode.com/problems/remove-covered-intervals/
// Given a list of intervals, remove all intervals that are covered by another interval in the list.
// Interval [a,b) is covered by interval [c,d) if and only if c <= a and b <= d.
// After doing so, return the number of remaining intervals.
func removeCoveredIntervals(intervals [][]int) int {

	// sort the slice based on first element -- if its a tie, we want
	// larger intervals to come first, such that if an element covers another element
	// the one that does the covering is always to the left
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	prev := intervals[0]
	i := 1
	covered := 0

	for i < len(intervals) {
		if covers(intervals[i], prev) {
			covered++
		} else {
			prev = intervals[i]
		}

		i++
	}

	return len(intervals) - covered

}

func main() {
	var test1 [][]int = [][]int{{1, 4}, {3, 6}, {2, 8}}

	fmt.Println(removeCoveredIntervals(test1)) // expects 2
}
