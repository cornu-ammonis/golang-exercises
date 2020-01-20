package main

import (
	"fmt"
)

func makeDynamic2DArray(n int) [][]bool {
	matrix := make([][]bool, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]bool, n)
	}

	return matrix
}

func buildAdjacencyMatrix(connections [][]int, n int) [][]bool {
	var matrix [][]bool = makeDynamic2DArray(n)

	for _, val := range connections {
		matrix[val[0]][val[1]] = true
	}

	return matrix
}

// https://leetcode.com/problems/number-of-operations-to-make-network-connected/
func makeConnected(n int, connections [][]int) int {
	return 0
}

func main() {
	fmt.Println("hello")

	test1 := [][]int{{0, 1}, {0, 2}, {1, 2}}
	testMatrix := buildAdjacencyMatrix(test1, 4)
	fmt.Println(testMatrix)
}
