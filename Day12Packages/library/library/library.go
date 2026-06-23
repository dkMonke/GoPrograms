// library.go — Library type and its methods, extracted into a reusable package.
// Demonstrates the package-level refactoring from the Day 08/12 single-file version.
// Sentinel errors and pointer-receiver methods (Add, Borrow, Return) provide
// a clean API that the main package imports and uses.
package library

import (
	"errors"
)

// Library is a collection of books exposed by the library package. Books holds
// the catalogue as a slice and is mutated in place by the Add, Borrow, and
// Return methods.
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
