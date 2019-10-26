package main

import (
	"fmt"

	"../ds/stack"
)

func main() {
	fmt.Printf("hello, world\n")

	testString := "abcdefg"

	s := make(stack.ByteStack, 0)
	s = s.Push(testString[0])
	fmt.Println(s)
	s = s.Push(testString[1])
	s = s.Push(testString[2])
	fmt.Println(s)
	fmt.Println(s.ToString())
	s, p := s.Pop()
	fmt.Println(p)
}
