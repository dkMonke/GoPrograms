// main.go — Day 12: Packages introduction.
// This file reuses the Person/Address struct-embedding example from Day 08
// to serve as a baseline before refactoring code into separate packages.
// Demonstrates struct embedding, value-receiver methods, and pointer parameters.
package main

import "fmt"

// Address holds the postal location fields (street, city, and zip) and is
// embedded into Person to demonstrate Go struct embedding.
type Address struct {
	Street, City, Zip string
}

// Person represents an individual with a Name, Age, and an embedded Address.
// Because Address is embedded, its fields (Street, City, Zip) are promoted and
// accessible directly on a Person value (e.g. p.City).
type Person struct {
	Name    string
	Age     int
	Address //embedded struct
}

// Describe returns a human-readable summary of the person, combining their
// name, age, street, and city. It uses a value receiver, so it operates on a
// copy and does not modify the original Person.
func (p Person) Describe() string {
	return fmt.Sprintf("%s (%d), lives at %s, %s", p.Name, p.Age, p.Street, p.City)

}

// birthday increments the person's Age by one. It takes a pointer so the
// mutation is reflected in the caller's Person value rather than a copy.
func birthday(p *Person) {
	p.Age++
}

// main builds a Person with an embedded Address, prints its description, ages
// it by one via birthday (passing a pointer), and shows promoted-field access
// through the embedded Address.
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
