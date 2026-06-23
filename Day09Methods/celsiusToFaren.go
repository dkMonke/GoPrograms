// celsiusToFaren.go — Day 09: Custom named types with conversion methods.
// Celsius and Farenheit are named float64 types. Each has a method to convert
// to the other unit. The program reads the target unit and value from CLI args,
// parses with strconv.ParseFloat, and prints the converted temperature.
package main

import (
	"fmt"
	"os"
	"strconv"
)

type Celsius float64
type Farenheit float64

func (c Celsius) ToFarenheit() Farenheit {
	return Farenheit(c)*9/5 + 32
}

func (f Farenheit) ToCelius() Celsius {
	return (Celsius(f) - 32) * 5 / 9
}

func main() {
	convertType := os.Args[1]
	convertValue, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Errorf("Encountered error", err)
	}
	if convertType == "Farenheit" {
		fmt.Printf("Converted value %d", Celsius(convertValue).ToFarenheit())
	} else {
		fmt.Printf("Converted value %d", Farenheit(convertValue).ToCelius())
	}

}
