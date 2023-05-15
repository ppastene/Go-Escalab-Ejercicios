package main

import "fmt"

type kelvin float64
type celsius float64
type farenheit float64

func kelvinToCelsius(k kelvin) celsius {

	// If this function returns k - 273.15 the compiler will throw an error
	// because the return type must be celsius, but instead
	// its returning the argument kelvin minus 273.15

	// This is correct because the return value is a celsius type
	return celsius(k - 273.15)
}

func main() {
	var k kelvin = 294.0
	var c celsius
	c = kelvinToCelsius(k)
	fmt.Println(c)

	// You cannot compare c and k because both are different types
	// so doing anything like c == k the compiler will throw an error
	// For example fmt.Println(c == k)

	// This is correct because is getting the type of each variable
	// and then compare then, so it will return false because celsius != kelvin
	fmt.Println(fmt.Sprintf("%T", k) == fmt.Sprintf("%T", c))
}
