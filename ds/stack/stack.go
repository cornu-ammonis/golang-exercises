package stack

type ByteStack []byte

func (s ByteStack) Push(v byte) ByteStack {
	return append(s, v)
}

func (s ByteStack) Pop() (ByteStack, byte) {
	l := len(s)

	return s[:l-1], s[l-1]
}
