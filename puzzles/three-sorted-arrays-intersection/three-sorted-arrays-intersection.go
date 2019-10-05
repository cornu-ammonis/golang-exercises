package main

import "fmt"

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	var iarr1, iarr2, iarr3 int
	var result []int

	for iarr1 < len(arr1) || iarr2 < len(arr2) || iarr3 < len(arr3) {
		if arr1[iarr1] == arr2[iarr2] && arr1[iarr1] == arr3[iarr3] {
			result = append(result, arr1[iarr1])
			iarr1 = iarr1 + 1
			iarr2 = iarr2 + 1
			iarr3 = iarr3 + 1
		} else if arr1[iarr1] < arr2[iarr2] {
			iarr1 = iarr1 + 1
		} else if arr1[iarr1] < arr3[iarr3] {
			iarr1 = iarr1 + 1
		} else if arr2[iarr2] < arr1[iarr1] || arr2[iarr2] < arr3[iarr3] {
			iarr2 = iarr2 + 1
		} else if arr3[iarr3] < arr1[iarr1] || arr3[iarr3] < arr2[iarr2] {
			iarr3 = iarr3 + 1
		}

	}
	return result
}

func main() {
	var iarr1 []int = make([]int, 5)
	iarr1[0] = 1
	r := arraysIntersection(iarr1, iarr1, iarr1)

	fmt.Printf("%v", r)
}
