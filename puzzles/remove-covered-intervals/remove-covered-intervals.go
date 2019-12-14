package main

import "fmt"

// https://leetcode.com/problems/remove-covered-intervals/
// Given a list of intervals, remove all intervals that are covered by another interval in the list.
// Interval [a,b) is covered by interval [c,d) if and only if c <= a and b <= d.
// After doing so, return the number of remaining intervals.
func removeCoveredIntervals(intervals [][]int) int {

}

func main() {
	var test1 [][]int = [][]int{{1, 4}, {3, 6}, {2, 8}}

	fmt.Println(removeCoveredIntervals(test1)) // expects 2
}
