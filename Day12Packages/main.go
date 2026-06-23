// main.go — Day 12: Packages introduction.
// This file reuses the Person/Address struct-embedding example from Day 08
// to serve as a baseline before refactoring code into separate packages.
// Demonstrates struct embedding, value-receiver methods, and pointer parameters.
package main

import "fmt"

type Address struct {
	Street, City, Zip string
}

type Person struct {
	Name    string
	Age     int
	Address //embedded struct
}

func (p Person) Describe() string {
	return fmt.Sprintf("%s (%d), lives at %s, %s", p.Name, p.Age, p.Street, p.City)

}

func birthday(p *Person) {
	p.Age++
}

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
