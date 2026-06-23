// assignment.go — Day 09 Assignment: Fixed-size Ring Buffer.
// RingBuffer stores up to 'size' integers; once full, new writes overwrite
// the oldest entry (circular behaviour). ReadAll returns the elements in
// insertion order by reading from head to end then wrapping around.
package main

import (
	"fmt"
)

// RingBuffer is a fixed-capacity circular buffer of integers. size is the maximum
// number of elements it will hold; head marks the oldest element's index used for
// overwrite and read ordering; data is the backing slice. tail is currently unused.
type RingBuffer struct {
	size int
	head int
	tail int
	data []int
}

// Write inserts v into the buffer. While the buffer is not yet full it appends v.
// Once the buffer reaches its size, it overwrites the oldest element at head and
// advances head, giving circular overwrite behaviour. It uses a pointer receiver
// to mutate the buffer in place.
func (r *RingBuffer) Write(v int) {
	if r.Len() == r.size { //size breached
		r.data[r.head] = v
		r.head++
	} else {
		r.data = append(r.data, v)
	}
}

// ReadAll returns a new slice containing the buffer's elements in insertion
// (oldest-to-newest) order. It does this by reading from head to the end of the
// backing slice and then wrapping around from the start up to head. The returned
// slice is a fresh copy and does not alias the buffer's internal storage.
func (r *RingBuffer) ReadAll() []int {
	var outputSlice []int
	outputSlice = append(outputSlice, r.data[r.head:]...)
	outputSlice = append(outputSlice, r.data[:r.head]...)
	return outputSlice

}

// Len returns the current number of elements stored in the buffer, which is the
// length of the backing slice (capped at size once the buffer is full).
func (r *RingBuffer) Len() int {
	return len(r.data)
}

// main is the entry point. It creates a RingBuffer of size 4, writes eight values
// so that the earliest entries are overwritten, then prints the ReadAll output and
// the buffer length to demonstrate the circular behaviour.
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
