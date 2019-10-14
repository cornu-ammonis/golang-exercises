// https://leetcode.com/problems/queens-that-can-attack-the-king/
package main

import "fmt"

func buildBoard(queens [][]int) [8][8]int {

}

func queensAttacktheKing(queens [][]int, king []int) [][]int {

}

// dy and dx represent the 'direction of travel' for x and y. for example,
// if dy = 0, dx = -1 then y remains unchanged and x decreaes by one each step.
// this particular example represents traversing left from the king.
// traversing the up-right diagonal would be represented by dy = 1, dx = 1
//
// returns the queen's coordinates, if found, and a flag representing whether we found a queen
func traverseDirection(board [8][8]int, kingY, kingX, dy, dx int) (pos []int, found bool) {

}

func main() {
	board := [8][8]int{} // best way to initialize an array like this!
	fmt.Println(board)
}
