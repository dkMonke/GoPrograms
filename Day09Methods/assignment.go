// assignment.go — Day 09 Assignment: Fixed-size Ring Buffer.
// RingBuffer stores up to 'size' integers; once full, new writes overwrite
// the oldest entry (circular behaviour). ReadAll returns the elements in
// insertion order by reading from head to end then wrapping around.
package main

import (
	"fmt"
)

type RingBuffer struct {
	size int
	head int
	tail int
	data []int
}

func (r *RingBuffer) Write(v int) {
	if r.Len() == r.size { //size breached
		r.data[r.head] = v
		r.head++
	} else {
		r.data = append(r.data, v)
	}
}

func (r *RingBuffer) ReadAll() []int {
	var outputSlice []int
	outputSlice = append(outputSlice, r.data[r.head:]...)
	outputSlice = append(outputSlice, r.data[:r.head]...)
	return outputSlice

}

func (r *RingBuffer) Len() int {
	return len(r.data)
}

func main() {
	var r = RingBuffer{size: 4}
	r.Write(1)
	r.Write(2)
	r.Write(3)
	r.Write(4)
	r.Write(5)
	r.Write(6)
	r.Write(7)
	r.Write(8)

	fmt.Println("Read all function output", r.ReadAll())
	fmt.Println("Length of the ring buffer", r.Len())

}
