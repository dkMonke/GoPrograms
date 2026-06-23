// assignment.go — Day 15 Assignment: Concurrent file reader.
// countLines reads a file's contents using os.ReadFile and prints the result.
// Each file is read in its own goroutine (go countLines(v)).
// Note: WaitGroup is commented out — without it, main may exit before
// goroutines finish, which is a known concurrency pitfall being explored.
package main

import (
	"fmt"
	"os"
	//"time"
)

func countLines(filename string) {
	//defer wg.Done()
	f, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error opening a file")
	}
	fmt.Printf("Filename %q length %q", filename, string(f))
}

func main() {
	fileNames := []string{"file1", "file2", "file3"}
	//var wg sync.WaitGroup
	for _, v := range fileNames {
		//wg.Add(1)
		go countLines(v)
	}
	//wg.Wait()

}
