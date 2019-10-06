package main

import "fmt"

type opt struct {
	row int
	col int
}

func options(row int, col int, grid [][]int) []opt {
	var opts []opt = make([]opt, 4)

	if row+1 < len(grid) {
		opts = append(opts, opt{row + 1, col})
	}

	if col+1 < len(grid[row]) {
		opts = append(opts, opt{row, col + 1})
	}

	if row > 0 {
		opts = append(opts, opt{row - 1, col})
	}

	if col > 0 {
		opts = append(opts, opt{row, col - 1})
	}

	return opts
}

func backtrack(row int, col int, grid [][]int, currentGold *int, maxGold *int) {
	var goldAtPos int = grid[row][col]

	if goldAtPos == 0 {
		return
	}

	*currentGold = *currentGold + goldAtPos

	if *currentGold > *maxGold {
		*maxGold = *currentGold
	}

	// clear out at this position to avoid duplicates
	grid[row][col] = 0

	for _, opt := range options(row, col, grid) {
		backtrack(opt.row, opt.col, grid, currentGold, maxGold)
	}

	*currentGold = *currentGold - goldAtPos // revert
	grid[row][col] = goldAtPos              // restore grid position as option for future backtracks

}

func getMaximumGold(grid [][]int) int {

	var currentGold int = 0
	var maxGold int = 0

	for row := range grid {
		for column := range grid[row] {
			backtrack(row, column, grid, &currentGold, &maxGold)
		}
	}

	return maxGold
}

func main() {
	var test1 [][]int = [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}
	fmt.Println("running first test, expect 24...")
	res1 := getMaximumGold(test1)
	fmt.Println(res1) // expect 24

	test2 := [][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}}
	fmt.Println("running second test, expect 28...")
	res2 := getMaximumGold(test2)
	fmt.Println(res2)
}
