package main

import (
	"fmt"
	"sort"
)

type Planets []string

func (p Planets) terraform() {
	for i := range p {
		p[i] = "New " + p[i]
	}
}

func main() {
	planets := []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	// Sort() is to order the elements of the slice in alphabetical order
	sort.StringSlice(planets).Sort()
	// Also you can use this to make it simple
	// sort.Strings(planets)
	fmt.Println(planets)

	// planets2 is a type of Planets which is a slice
	planets2 := Planets{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	// And the Planets type holds a method to modify the slice
	// Keep in mind, the mothod is not returning anything but modifying the slice outside of it
	planets2.terraform()
	fmt.Println(planets2)

	// append() allows to insert new elements into the slice
	// Like with Println() is a variadic function so it can accept one or more elements
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs = append(dwarfs, "Orcus", "Salacia")
	fmt.Println(dwarfs)
}
