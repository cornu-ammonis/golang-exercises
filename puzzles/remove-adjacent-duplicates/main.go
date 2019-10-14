package main

import (
	"fmt"

	"../../ds/stack"
)

func singlePassRemove(s string, k int) string {

}

func removeDuplicates(s string, k int) string {
	var prevLen int = 0
	var currentLen int = len(s)

	for prevLen != currentLen {
		prevLen = currentLen

		s = singlePassRemove(s, k)
	}
}

func main() {
	s := make(stack.ByteStack, 0)
	s = s.Push('a')
	fmt.Println(s)
}
