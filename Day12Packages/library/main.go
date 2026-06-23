// main.go — Entry point for the library application.
// Imports the local library package and exercises its API: creating a Book,
// adding it to a Library, borrowing, and returning. Shows how to use
// types and methods defined in a separate package via the import path.
package main

import (
	"fmt"
	"library/library"
)

func main() {
	Book1 := library.Book{
		Title:     "TestBook",
		Author:    "Dinesh",
		Pages:     145,
		Available: true,
	}
	lib1 := library.Library{}
	lib1.Add(Book1)
	fmt.Println(lib1.Books)
	err := lib1.Borrow("TestBook")
	fmt.Println(lib1.Books)
	fmt.Println("Error1", err)
	err1 := lib1.Return("TestBook")
	if err1 != nil {
		fmt.Println("Error2", err1)
	}

}
