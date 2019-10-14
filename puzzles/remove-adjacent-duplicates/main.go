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
	sameCount := 0
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
			sameCount = 0
			prev = c
		}
	}

	return stk.toString()
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
	s := make(stack.ByteStack, 0)
	s = s.Push('a')
	fmt.Println(s)
}
