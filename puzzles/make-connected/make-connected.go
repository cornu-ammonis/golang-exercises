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

	}

}

// https://leetcode.com/problems/number-of-operations-to-make-network-connected/
func makeConnected(n int, connections [][]int) int {
	return 0
}

func main() {
	fmt.Println("hello")
}
