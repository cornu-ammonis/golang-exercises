// exercise taken from https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/

package main

import (
	"fmt"
)

type ch struct {
	char  byte
	count int
}

// solution is O(n) and benchmarks at 0ms runtime and 5.8 MB memory usage
// my first attempt solution, found commented at the bottom of the file
// benchmarked at 236ms and 8mb memory
func removeDuplicates(s string, k int) string {
	stack := make([]ch, 0)
	firstCh := ch{char: s[0], count: 1}
	stack = append(stack, firstCh)

	for i := 1; i < len(s); i++ {
		c := s[i]

		var prev ch
		if len(stack) > 0 {
			prev = stack[len(stack)-1]
		}

		if (prev != ch{}) && (prev.char == c) {
			stack = append(stack, ch{
				char:  c,
				count: 1 + prev.count,
			})
		} else {
			stack = append(stack, ch{
				char:  c,
				count: 1,
			})
		}

		if stack[len(stack)-1].count == k {
			for i := 0; i < k; i++ {
				stack = stack[:len(stack)-1]
			}
		}

	}

	res := make([]byte, 0)

	for i := 0; i < len(stack); i++ {
		res = append(res, stack[i].char)
	}
	return string(res)
}

func main() {
	test1 := "abcd"
	k1 := 2

	test2 := "deeedbbcccbdaa"
	k2 := 3

	test3 := "pbbcggttciiippooaais"
	k3 := 2

	fmt.Println(removeDuplicates(test1, k1)) // expect "abcd"
	fmt.Println(removeDuplicates(test2, k2)) // expect "aa"
	fmt.Println(removeDuplicates(test3, k3)) // expect "ps"
}

/*
OLD SOLUTION. not O(n)

type ByteStack []byte

func (s ByteStack) Push(v byte) ByteStack {
	return append(s, v)
}

func (s ByteStack) Pop() (ByteStack, byte) {
	l := len(s)

	return s[:l-1], s[l-1]
}

func (s ByteStack) ToString() string {
	str := ""
	for _, c := range s {
		str = str + string(c)
	}

	return str
}

func removeKFromStack(stk ByteStack, k int) ByteStack {
	for i := 0; i < k; i = i + 1 {
		stk, _ = stk.Pop()
	}

	return stk
}

func singlePassRemove(s string, k int) string {
	prev := byte(0)
	sameCount := 0 // starts at 0 because we have 0 same char in a row
	stk := make(ByteStack, 0)
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
*/
