package main

import "fmt"

// actually not needed, just expiramenting
func safeIncrement(index int, array []int) int {
	if index < len(array)-1 {
		return index + 1
	}
	return index
}

// not needing, expirementing
func safeIncrementPointer(index *int, array []int) {
	if *index < len(array) {
		*index = *index + 1
	}
}

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	var iarr1, iarr2, iarr3 int
	var result []int

	// we use && because once any index gets past the last element, we can no longer discover
	// overlapping elements
	for iarr1 < len(arr1) && iarr2 < len(arr2) && iarr3 < len(arr3) {
		if arr1[iarr1] == arr2[iarr2] && arr1[iarr1] == arr3[iarr3] {
			result = append(result, arr1[iarr1])

			iarr1++
			iarr2++
			iarr3++
		} else if arr1[iarr1] < arr2[iarr2] || arr1[iarr1] < arr3[iarr3] {
			iarr1++
		} else if arr2[iarr2] < arr1[iarr1] || arr2[iarr2] < arr3[iarr3] {
			iarr2++
		} else if arr3[iarr3] < arr1[iarr1] || arr3[iarr3] < arr2[iarr2] {
			iarr3++
		}
	}
	return result
}

func main() {

	arr1 := []int{2, 3, 5, 7, 11, 13}
	arr2 := []int{2, 4, 5, 8, 9, 13}
	arr3 := []int{2, 3, 4, 5, 7, 11, 13}

	r := arraysIntersection(arr1, arr2, arr3)

	fmt.Printf("%v", r)

	var i int = 4
	var arr []int = make([]int, 8)
	fmt.Println(i)
	safeIncrementPointer(&i, arr)
	fmt.Println(i)

	for i < len(arr) {
		safeIncrementPointer(&i, arr)
	}
	fmt.Println(i)
}
