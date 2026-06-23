// main.go — Day 09: Methods on types — Stack implementation.
// Stack wraps a []int slice. Pointer-receiver methods Push, Pop, and Len
// provide classic LIFO operations. Pop returns (value, ok) using the
// comma-ok pattern to handle empty-stack gracefully.
package main

import "fmt"

// Stack is a last-in-first-out (LIFO) collection of integers backed by a slice.
// The items field holds the elements with the top of the stack at the end.
type Stack struct {
	items []int
}

// Push adds v to the top of the stack by appending it to the items slice. It
// uses a pointer receiver so the underlying slice is updated in place.
func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

// Pop removes and returns the element at the top of the stack along with true.
// If the stack is empty it returns (0, false), following the comma-ok pattern so
// callers can distinguish an empty stack from a zero value.
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	last := len(s.items) - 1
	v := s.items[last]
	s.items = s.items[:last]
	return v, true
}

// Len returns the number of elements currently on the stack.
func (s *Stack) Len() int {
	return len(s.items)

}

// main is the entry point. It pushes three values onto a stack, prints the
// length, then repeatedly pops until empty, printing each popped value to
// demonstrate LIFO ordering.
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
