package main

/*
	taken from https://leetcode.com/problems/cells-with-odd-values-in-a-matrix/
	Given n and m which are the dimensions of a matrix initialized by zeros and given an
	array indices where indices[i] = [ri, ci]. For each pair of [ri, ci] you have to increment
	all cells in row ri and column ci by 1.

	Return the number of cells with odd values in the matrix after applying the increment to all indices.
*/

func oddCells(n int, m int, indices [][]int) int {
	var rows []int = make([]int, n)
	var cols []int = make([]int, m)
	var result int = 0

	for _, entry := range indices {
		var ri int = entry[0]
		var ci int = entry[1]

		rows[ri]++
		cols[ci]++
	}

	for _, row := range rows {
		for _, col := range cols {
			if isOdd(row + col) {
				result++
			}
		}
	}

	return result
}

func isOdd(n int) bool {
	return n%2 == 1
}
