package main

import (
	"fmt"
)

func makeDynamic2DArray(n int) [][]int {
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	return matrix
}

func buildAdjacencyMatrix(connections [][]int, n int) [][]int {
	var matrix [][]int = makeDynamic2DArray(n)

}

// https://leetcode.com/problems/number-of-operations-to-make-network-connected/
func makeConnected(n int, connections [][]int) int {
	return 0
}

func main() {
	fmt.Println("hello")
}
