package main

import (
	"fmt"

	"../../ds/stack"
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
		// assumes undirected graph, so the matrix is strictly symmetric
		matrix[val[0]][val[1]] = true
		matrix[val[1]][val[0]] = true
	}

	return matrix
}

// https://leetcode.com/problems/number-of-operations-to-make-network-connected/
func makeConnected(n int, connections [][]int) int {

	if len(connections) < n-1 {
		return -1
	}

	var matrix [][]bool = buildAdjacencyMatrix(connections, n)

	var seenVertex []bool = make([]bool, n)
	var seenVertexCount int = 0

	var components int = 0
	var i int = 0
	stk := make(stack.IntStack, 0)

	for seenVertexCount < n {
		// skip starting DFS from vertices we've already seen
		if seenVertex[i] {
			i++
			continue
		}

		// if we havent seen this before, we are at a new component
		components++
		seenVertexCount++ // we are at the vertex i so make sure to record it
		seenVertex[i] = true

		for vrtx := 0; vrtx < n; vrtx++ {
			if matrix[i][vrtx] && !seenVertex[vrtx] {
				stk = stk.Push(vrtx)
				seenVertex[vrtx] = true
				seenVertexCount++
			}
		}

		var stackCounter int = len(stk)
		var vrtx int = -1
		for len(stk) > 0 {
			stk, vrtx = stk.Pop()
			stackCounter--

			for v := 0; v < n; v++ {
				if matrix[vrtx][v] == true && !seenVertex[v] {
					stk = stk.Push(v)
					stackCounter++
					seenVertex[v] = true
					seenVertexCount++
				}
			}
		}

		i++

	}

	return components - 1
}

func main() {
	fmt.Println("hello")

	test1 := [][]int{{0, 1}, {0, 2}, {1, 2}}
	testMatrix := buildAdjacencyMatrix(test1, 4)
	fmt.Println(testMatrix)
	fmt.Println(makeConnected(4, test1)) // expect 1

	test2 := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}
	fmt.Println(makeConnected(6, test2)) // expect 2

	test3 := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}
	fmt.Println(makeConnected(6, test3)) // expect -1

	test4 := [][]int{{0, 1}, {0, 2}, {3, 4}, {2, 3}}
	fmt.Println(makeConnected(5, test4)) // expect 0
}
