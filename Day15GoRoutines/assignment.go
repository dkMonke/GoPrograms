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

// countLines reads the entire contents of the file named by filename using
// os.ReadFile and prints the filename alongside its contents. If the file
// cannot be opened an error message is printed, but execution continues and
// the (empty) contents are still printed. This function is intended to be run
// as a goroutine, one per file, so multiple files are read concurrently.
func countLines(filename string) {
	//defer wg.Done()
	f, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error opening a file")
	}
	fmt.Printf("Filename %q length %q", filename, string(f))
}

// main builds a slice of file names and launches a countLines goroutine for
// each one so they are read concurrently. The sync.WaitGroup is intentionally
// commented out, so main may return (and the program exit) before the
// goroutines finish — a deliberate demonstration of an unsynchronised
// concurrency pitfall.
func main() {
	fileNames := []string{"file1", "file2", "file3"}
	//var wg sync.WaitGroup
	for _, v := range fileNames {
		//wg.Add(1)
		go countLines(v)
	}
	//wg.Wait()

}
