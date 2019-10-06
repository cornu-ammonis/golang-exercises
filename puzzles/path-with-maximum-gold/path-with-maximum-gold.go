package main

type opt struct {
	row int
	col int
}

func options(row int, column int, grid [][]int) []opt {

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

}
