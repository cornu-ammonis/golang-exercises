package main

import "fmt"

func start(slot []int) int {
	return slot[0]
}

func end(slot []int) int {
	return slot[1]
}

func maxStartSlot(slot1 []int, slot2 []int) []int {
	if slot1[0] > slot2[0] {
		return slot1
	}
	return slot2
}

func maxStart(slot1 []int, slot2 []int) int {
	return start(maxStartSlot(slot1, slot2))
}

func minEndSlot(slot1 []int, slot2 []int) []int {
	if slot1[1] < slot2[1] {
		return slot1
	}
	return slot2
}

func minEnd(slot1 []int, slot2 []int) int {
	return end(minEndSlot(slot1, slot2))
}

//https://leetcode.com/problems/meeting-scheduler/
func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	res := make([]int, 2)
	i, j := 0, 0

	for i < len(slots1) && j < len(slots2) {
		slot1 := slots1[i]
		slot2 := slots2[j]

	}
}

func main() {
	fmt.Println("hi")
}
