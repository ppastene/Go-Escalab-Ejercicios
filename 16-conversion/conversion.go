package main

import (
	"fmt"
	"strconv"
)

// Golng is very strict with the data types, you cannot do a math with an integer and a float
// So the code must convert the values to do something

func main() {
	// Since age have an integer value it must be converted to float64 to be used on marsAge
	age := 41
	marsAge := float64(age)
	marsDays := 687.0
	earthDays := 365.2425
	marsAge = marsAge * earthDays / marsDays
	fmt.Println("Mars have", marsAge, "years")

	// In Golang there are functions to convert any type of data
	// Here you can convert a float64 to integer directly into the Println function
	fmt.Println("The earth have", int(earthDays), "days per year")

	// And also not only works between numbers,
	// here you can convert a numeric type into a string with strconv.Itoa() and concat directly into the string
	earthHours := int(earthDays) * 24
	fmt.Println("Per year the earth have "+strconv.Itoa(earthHours), "hours")
}
