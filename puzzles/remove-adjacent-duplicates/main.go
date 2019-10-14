// exercise taken from https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/

package main

import (
	"fmt"
)

type ch struct {
	char  byte
	count int
}

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
