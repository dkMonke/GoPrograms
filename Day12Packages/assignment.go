// assignment.go — Day 12 Assignment: Library system (single-file version).
// This is the monolithic version before being split into the library/ package.
// All types (Book, Library) and methods (Add, Borrow, Return) live in main,
// serving as a before/after comparison for the package refactoring exercise.
package main

import (
	"errors"
	"fmt"
)

// Book represents a single book in the library catalogue. Title and Author
// describe the work, Pages records its length, and Available indicates whether
// the book is currently on the shelf (true) or borrowed (false).
type Book struct {
	Title, Author string
	Pages         int
	Available     bool
}

// Library is a collection of books. Books holds the catalogue as a slice and
// is mutated in place by the Add, Borrow, and Return methods.
type Library struct {
	Books []Book
}

// Add appends the given Book to the library's catalogue. It uses a pointer
// receiver so the underlying Books slice on the original Library is updated.
func (l *Library) Add(b Book) {
	l.Books = append(l.Books, b)
}

// BookNotFoundErr is the sentinel error returned when no book matches the
// requested title during a Borrow or Return operation.
var BookNotFoundErr = errors.New("The Book requested is not found")

// ErrBookReturned is the sentinel error returned by Return when the book is
// already marked Available, indicating it was never borrowed (or already returned).
var ErrBookReturned = errors.New("The book has been already returned")

// Borrow marks the first book matching Title as unavailable. It returns nil on
// success or BookNotFoundErr if no book with that title exists in the library.
func (l *Library) Borrow(Title string) error {
	for i := range l.Books {
		if l.Books[i].Title == Title {
			l.Books[i].Available = false
			return nil
		}
	}
	return BookNotFoundErr
}

// Return marks the first book matching Title as available again. It returns
// BookNotFoundErr if no such book exists, ErrBookReturned if the book is already
// available (so cannot be returned), or nil on a successful return.
func (l *Library) Return(Title string) error {
	for i := range l.Books {
		if l.Books[i].Title == Title {
			if l.Books[i].Available == true {
				return ErrBookReturned
			}
			l.Books[i].Available = true
			return nil
		}
	}
	return BookNotFoundErr

}

// main demonstrates the single-file library workflow end to end: it creates a
// Book, adds it to a Library, borrows it, returns it once successfully, then
// attempts a second return to show the ErrBookReturned error path, printing
// the catalogue state and any errors after each step.
func main() {
	Book1 := Book{
		Title:     "Moby dick",
		Author:    "Jefferson",
		Pages:     100,
		Available: true,
	}

	lib1 := Library{}
	lib1.Add(Book1)
	fmt.Println(lib1.Books)
	err := lib1.Borrow("Moby dick")
	fmt.Println(lib1.Books)
	fmt.Println("Error1", err)
	err1 := lib1.Return("Moby dick")
	if err1 != nil {
		fmt.Println("Error2", err1)
	}
	fmt.Println(lib1.Books)
	err2 := lib1.Return("Moby dick")
	if err2 != nil {
		fmt.Println("Erro3", err2)
	}
	fmt.Println(lib1.Books)

}
