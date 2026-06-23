// main.go — Day 09: Methods on types — Stack implementation.
// Stack wraps a []int slice. Pointer-receiver methods Push, Pop, and Len
// provide classic LIFO operations. Pop returns (value, ok) using the
// comma-ok pattern to handle empty-stack gracefully.
package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	last := len(s.items) - 1
	v := s.items[last]
	s.items = s.items[:last]
	return v, true
}

func (s *Stack) Len() int {
	return len(s.items)

}

func main() {
	var s Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println("length:", s.Len())

	for s.Len() > 0 {
		v, _ := s.Pop()
		fmt.Println("popped:", v)
	}
}
