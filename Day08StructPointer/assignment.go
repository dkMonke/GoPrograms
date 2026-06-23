// assignment.go — Day 08 Assignment: Simple Library management system.
// Book and Library structs model a book catalogue. Pointer-receiver methods
// (Add, Borrow, Return) mutate the Library's Books slice in place.
// Uses sentinel errors (BookNotFoundErr, ErrBookReturned) to signal domain failures.
package main

import (
	"errors"
	"fmt"
)

type Book struct {
	Title, Author string
	Pages         int
	Available     bool
}

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
