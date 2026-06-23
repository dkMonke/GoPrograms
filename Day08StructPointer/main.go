// main.go — Day 08: Structs and Pointers.
// Defines Address and Person structs with struct embedding (Person embeds Address).
// Demonstrates a value-receiver method (Describe), a pointer parameter (birthday)
// that mutates the struct in place, and promoted field access via embedding (me.City).
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
