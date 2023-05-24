package main

import (
	"fmt"
	"strings"
)

func main() {
	// A slice is not an array but a 'view', 'window' or 'pointer' to it
	// This is an array because is defined in one way or another its size
	// array := [...]string{"one","two","three", "four", "five"}
	// array := [3]string{"one","two","three"}
	// This is a slice because doesnt have a defined size so its initialized along with an array
	// slice := []string{"one","two","three", "four", "five"}

	// Slicing is shown as an 'open-half range', For example planets[0:4] will start from index 0 but
	// will count and stop at the forth element, in this array at the [3] index
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	terrestial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]
	fmt.Println("Slices:", terrestial, gasGiants, iceGiants)

	// You can assign the slice to a variable and modify their values either on the new or previous slice
	// Despite of being 2 slices even with different values both are still looking to the same array since it cannot be modified
	iceGiantsMarkII := iceGiants
	iceGiants[1] = "Poseidon"
	fmt.Println("Slice:", iceGiants, "| New slice from iceGiants:", iceGiantsMarkII)

	// If you skip the first value of the slice, it will begin counting up to the amount of elements
	fmt.Println("Slice from index 0 counting 4 elements:", planets[:4])
	// If you skip the second value of the slice, it will begin counting from that index
	fmt.Println("Slice from index 6 counting the rest of the elements", planets[6:])
	// If you skip both values it will count all the slice
	fmt.Println("Slice from start to end:", planets[:])

	// In Golang many functions works with slices instead of arrays.
	// So, if you need to show every element of an slice the best way is to create an array and then the slice
	dwarfArray := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfSlice := dwarfArray[:]
	// Also is possible to create an array and slice at the same time
	// Drspite of being a slice, the compiler will create its array
	dwarfsSlice2 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Printf("dwarfArray is a: %T | dwarfSlice is a: %T | dwarfsSlice2 is a: %T\n", dwarfArray, dwarfSlice, dwarfsSlice2)

	// In hyperspace is argument is a copy of the slice but despite of not return anything
	// it modifies the values of the slice outside of the function,
	// that's because any modification of the slice will also modify it correspondant array
	planetsOnSpace := []string{"     Venus   ", "      Earth    ", "    Mars"}
	fmt.Println(planetsOnSpace)
	hyperspace(planetsOnSpace)
	fmt.Println(strings.Join(planetsOnSpace, "-"))
}

// Function to delete spaces from the slice elements
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}
