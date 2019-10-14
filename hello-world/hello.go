package main

import (
	"fmt"

	"../ds/stack"
)

func main() {
	fmt.Printf("hello, world\n")

	s := make(stack.ByteStack, 0)
	s = s.Push(1)
	fmt.Println(s)
	s = s.Push(2)
	s = s.Push(3)
	fmt.Println(s)
	s, p := s.Pop()
	fmt.Println(p)
}
