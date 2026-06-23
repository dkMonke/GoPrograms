// library.go — Library type and its methods, extracted into a reusable package.
// Demonstrates the package-level refactoring from the Day 08/12 single-file version.
// Sentinel errors and pointer-receiver methods (Add, Borrow, Return) provide
// a clean API that the main package imports and uses.
package library

import (
	"errors"
)

type Library struct {
	Books []Book
}

func (l *Library) Add(b Book) {
	l.Books = append(l.Books, b)
}

var BookNotFoundErr = errors.New("The Book requested is not found")
var ErrBookReturned = errors.New("The book has been already returned")

func (l *Library) Borrow(Title string) error {
	for i := range l.Books {
		if l.Books[i].Title == Title {
			l.Books[i].Available = false
			return nil
		}
	}
	return BookNotFoundErr
}

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
