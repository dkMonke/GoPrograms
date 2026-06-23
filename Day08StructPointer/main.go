// main.go — Day 08: Structs and Pointers.
// Defines Address and Person structs with struct embedding (Person embeds Address).
// Demonstrates a value-receiver method (Describe), a pointer parameter (birthday)
// that mutates the struct in place, and promoted field access via embedding (me.City).
package main

import "fmt"

// Address holds the postal components of a location: street, city and zip code.
type Address struct {
	Street, City, Zip string
}

// Person represents an individual with a Name and Age. It embeds Address, so an
// address's fields (Street, City, Zip) are promoted and accessible directly on a
// Person value.
type Person struct {
	Name    string
	Age     int
	Address //embedded struct
}

// Describe returns a human-readable summary of the person, combining their name,
// age and address. It uses a value receiver, so it operates on a copy and does
// not modify the original Person.
func (p Person) Describe() string {
	return fmt.Sprintf("%s (%d), lives at %s, %s", p.Name, p.Age, p.Street, p.City)

}

// birthday increments the age of the person pointed to by p. Because it takes a
// pointer, the change is visible to the caller after the function returns.
func birthday(p *Person) {
	p.Age++
}

// main is the entry point. It constructs a Person with an embedded Address,
// prints the Describe output, ages the person via the birthday pointer function,
// and demonstrates promoted-field access (me.City) through embedding.
func main() {
	me := Person{
		Name: "Alice",
		Age:  29,
		Address: Address{
			Street: "1 Main St",
			City:   "Anytown",
			Zip:    "12345",
		},
	}

	fmt.Println(me.Describe())
	birthday(&me)
	fmt.Println("After birthday:", me.Age)
	fmt.Println("Direct field access via embedding:", me.City)

}
