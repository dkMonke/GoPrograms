// book.go — Defines the Book struct for the library package.
// All fields are exported (capitalised) so they can be accessed from main.
// Separating the data model into its own file keeps the package organised.
package library

// Book represents a single book in the library catalogue. Title and Author
// describe the work, Pages records its length, and Available indicates whether
// the book is currently on the shelf (true) or borrowed (false). All fields are
// exported so they can be set and read from the importing main package.
type Book struct {
	Title, Author string
	Pages         int
	Available     bool
}
