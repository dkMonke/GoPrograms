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

// Celsius is a named float64 type representing a temperature in degrees Celsius.
type Celsius float64

// Farenheit is a named float64 type representing a temperature in degrees Fahrenheit.
type Farenheit float64

// ToFarenheit converts the Celsius value to its Fahrenheit equivalent using the
// formula (c * 9/5) + 32 and returns the result as a Farenheit value.
func (c Celsius) ToFarenheit() Farenheit {
	return Farenheit(c)*9/5 + 32
}

// ToCelius converts the Farenheit value to its Celsius equivalent using the
// formula (f - 32) * 5/9 and returns the result as a Celsius value.
func (f Farenheit) ToCelius() Celsius {
	return (Celsius(f) - 32) * 5 / 9
}

// main is the entry point. It reads a target unit and a numeric value from the
// command-line arguments (os.Args[1] and os.Args[2]), parses the value with
// strconv.ParseFloat, and prints the converted temperature. A target of
// "Farenheit" treats the input as Celsius and converts to Fahrenheit; any other
// value treats the input as Fahrenheit and converts to Celsius.
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
