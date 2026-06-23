// assignment.go — Day 08 Assignment: Simple Library management system.
// Book and Library structs model a book catalogue. Pointer-receiver methods
// (Add, Borrow, Return) mutate the Library's Books slice in place.
// Uses sentinel errors (BookNotFoundErr, ErrBookReturned) to signal domain failures.
package main

import (
	"errors"
	"fmt"
)

// Book represents a single catalogue entry. Title and Author identify the book,
// Pages records its length, and Available indicates whether it is currently on
// the shelf (true) or borrowed (false).
type Book struct {
	Title, Author string
	Pages         int
	Available     bool
}

// Library is a collection of books. Books holds the catalogue as a slice and is
// mutated in place by the pointer-receiver methods Add, Borrow and Return.
type Library struct {
	Books []Book
}

// Add appends the given book to the library's catalogue. It uses a pointer
// receiver so the append reassigns the receiver's Books slice in place.
func (l *Library) Add(b Book) {
	l.Books = append(l.Books, b)
}

// BookNotFoundErr is the sentinel error returned when no book matches the
// requested title during a Borrow or Return operation.
var BookNotFoundErr = errors.New("The Book requested is not found")

// ErrBookReturned is the sentinel error returned by Return when the matched book
// is already marked available, indicating it was never borrowed (or returned twice).
var ErrBookReturned = errors.New("The book has been already returned")

// Borrow marks the first book matching Title as unavailable and returns nil on
// success. If no book with that title exists it returns BookNotFoundErr. It does
// not check whether the book was already borrowed.
func (l *Library) Borrow(Title string) error {
	for i := range l.Books {
		if l.Books[i].Title == Title {
			l.Books[i].Available = false
			return nil
		}
	}
	return BookNotFoundErr
}

// Return marks the first book matching Title as available again and returns nil
// on success. If the matched book is already available it returns ErrBookReturned;
// if no book with that title exists it returns BookNotFoundErr.
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

// main is the entry point. It creates a library, adds one book, then exercises
// the Borrow and Return flow — including the duplicate-return case — printing the
// catalogue state and any returned errors after each step.
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
