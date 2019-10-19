package main

import (
	"fmt"
	"sort"
)

func start(slot []int) int {
	return slot[0]
}

func end(slot []int) int {
	return slot[1]
}

func maxStart(slot1 []int, slot2 []int) int {
	if start(slot1) > start(slot2) {
		return start(slot1)
	}
	return start(slot2)
}

func minEnd(slot1 []int, slot2 []int) int {
	if end(slot1) < end(slot2) {
		return end(slot1)
	}
	return end(slot2)
}

func incrementMinEndIndex(slot1 []int, slot2 []int, i *int, j *int) {
	if end(slot1) < end(slot2) {
		*i = *i + 1
	} else {
		*j = *j + 1
	}
}

func sortSlots(slots [][]int) {
	sort.Slice(slots, func(i, j int) bool {
		return slots[i][0] < slots[j][0]
	})
}

//https://leetcode.com/problems/meeting-scheduler/
func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	sortSlots(slots1)
	sortSlots(slots2)

	res := make([]int, 0)
	i, j := 0, 0

	for i < len(slots1) && j < len(slots2) {
		slot1 := slots1[i]
		slot2 := slots2[j]

		if (minEnd(slot1, slot2) - maxStart(slot1, slot2)) >= duration {
			return []int{maxStart(slot1, slot2), maxStart(slot1, slot2) + duration}
		}

		incrementMinEndIndex(slot1, slot2, &i, &j)
	}

	return res
}

func main() {
	fmt.Println("hi")

	test1 := [][]int{{10, 50}, {60, 120}, {140, 210}}
	test2 := [][]int{{10, 15}, {60, 70}}
	duration1 := 8

	unsortedTest := [][]int{{60, 120}, {10, 50}, {140, 210}}
	fmt.Println(unsortedTest)
	sortSlots(unsortedTest)
	fmt.Println(unsortedTest)

	fmt.Println(minAvailableDuration(test1, test2, duration1))
}
