package main

import (
	"fmt"

	"../../ds/stack"
)

func removeKFromStack(stk stack.ByteStack, k int) stack.ByteStack {
	for i := 0; i < k; i = i + 1 {
		stk, _ = stk.Pop()
	}

	return stk
}

func singlePassRemove(s string, k int) string {
	prev := byte(0)
	sameCount := 0 // starts at 0 because we have 0 same char in a row
	stk := make(stack.ByteStack, 0)
	for i, _ := range s {
		var c byte = s[i]
		stk = stk.Push(c)

		if c == prev {
			sameCount = sameCount + 1
			if sameCount == k {
				stk = removeKFromStack(stk, k)
			}
		} else {
			sameCount = 1 // starts at one here because we have 1 same char in a row
			prev = c
		}
	}

	return stk.ToString()
}

func removeDuplicates(s string, k int) string {
	var prevLen int = 0
	var currentLen int = len(s)

	for prevLen != currentLen {
		prevLen = currentLen

		s = singlePassRemove(s, k)
		currentLen = len(s)
	}

	return s
}

func main() {
	test1 := "abcd"
	k1 := 2

	test2 := "deeedbbcccbdaa"
	k2 := 3

	fmt.Println(removeDuplicates(test1, k1))
	fmt.Println(removeDuplicates(test2, k2))
}
