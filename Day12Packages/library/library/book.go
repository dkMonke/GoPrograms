// book.go — Defines the Book struct for the library package.
// All fields are exported (capitalised) so they can be accessed from main.
// Separating the data model into its own file keeps the package organised.
package library

type Book struct {
	Title, Author string
	Pages         int
	Available     bool
}
