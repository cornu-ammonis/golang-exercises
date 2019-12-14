package main

import (
	"fmt"
	"sort"
)

// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func remove(slice [][]int, s int) [][]int {
	return append(slice[:s], slice[s+1:]...)
}

// https://leetcode.com/problems/remove-covered-intervals/
// Given a list of intervals, remove all intervals that are covered by another interval in the list.
// Interval [a,b) is covered by interval [c,d) if and only if c <= a and b <= d.
// After doing so, return the number of remaining intervals.
func removeCoveredIntervals(intervals [][]int) int {

	// sort the slice based on first element for now
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	fmt.Println(intervals)
	intervals = remove(intervals, 1)
	fmt.Println(intervals)
	return 0
}

func main() {
	var test1 [][]int = [][]int{{1, 4}, {3, 6}, {2, 8}}

	fmt.Println(removeCoveredIntervals(test1)) // expects 2
}
