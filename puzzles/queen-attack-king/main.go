// https://leetcode.com/problems/queens-that-can-attack-the-king/
package main

import "fmt"

func buildBoard(queens [][]int) [8][8]bool {
	board := [8][8]bool{}

	for _, v := range queens {
		board[v[0]][v[1]] = true
	}

	return board
}

// dy and dx represent the 'direction of travel' for x and y. for example,
// if dy = 0, dx = -1 then y remains unchanged and x decreaes by one each step.
// this particular example represents traversing left from the king.
// traversing the up-right diagonal would be represented by dy = 1, dx = 1
//
// returns the queen's coordinates, if found, and a flag representing whether we found a queen
func traverseDirection(board [8][8]bool, kingX, kingY, dx, dy int) (coords []int, found bool) {

	x := kingX
	y := kingY
	nextX := kingX + dx
	nextY := kingY + dy

	for nextX >= 0 && nextX < 8 && nextY >= 0 && nextY < 8 {
		x = nextX
		y = nextY
		nextX = x + dx
		nextY = y + dy

		if board[x][y] {
			return []int{x, y}, true
		}
	}

	return []int{}, false
}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	res := make([][]int, 0)
	board := buildBoard(queens)
	kingX := king[0]
	kingY := king[1]

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			// this would represent no movement, so we skip it
			if dy == 0 && dx == 0 {
				continue
			}

			coords, found := traverseDirection(board, kingX, kingY, dx, dy)
			if found {
				res = append(res, coords)
			}
		}
	}

	return res
}

func main() {
	board := [8][8]int{} // best way to initialize an array like this!
	fmt.Println(board)
}
