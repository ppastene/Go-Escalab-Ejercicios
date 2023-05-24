package main

import "fmt"

func main() {
	// An array is a collection of data
	// In Golang the array must be of one type only, it cannot hold a string and an integer in the same collection
	// Here the variable holds an array assigned with 9 string elements, and 8 of those have a value
	var planets [9]string
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"
	planets[3] = "Mars"
	planets[4] = "Jupiter"
	planets[5] = "Saturn"
	planets[6] = "Uranus"
	planets[7] = "Neptune"
	// To access to the value mus be accessed with its index value
	// The first value counts from 0
	earth := planets[2]
	fmt.Println(earth)
	// It have a fixed size of elements, and despite of not assign anything it will show a default value
	// In the case of string, it will show ""
	fmt.Println(len(planets))
	fmt.Println(planets[8] == "")
	// You cannot show an index outside of the size of the array, so the following is wrong
	// fmt.Println(planets[9] == "")
	// With compiste literals is possible to create an array and assign values at the same time
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	// And you can create an array which the compiler will detect its size
	galaxies := [...]string{
		"Andromeda",
		"Black Eye",
		"Condor",
		"Hockey Stick",
		"Little Sombrero",
		"Medusa Merger",
		"Milky Way",
		"Peekaboo",
		"Pinwheel",
		"Sunflower",
		"Whirlpool",
	}
	// How to iterate an array, with the for loop
	for i := 0; i < len(dwarfs); i++ {
		fmt.Println(i, dwarfs[i])
	}
	for i, planet := range planets {
		fmt.Println(i, planet)
	}
	// The index can be skipped using an underscore
	for _, galaxy := range galaxies {
		fmt.Println(galaxy)
	}

	// How to copy an array
	planetMarkII := planets
	planets[2] = "toasty!"
	// This is the original array with the value of index 2 modified
	fmt.Println(planets)
	// This is the copy of the array which the modification is not present
	fmt.Println(planetMarkII)

	// In this function holds a copy of the array so is managed by value and not by reference
	// It returns the array keeping its size with all the values modified
	planets2 := modifyArray(planets)
	fmt.Println(planets2)

	// And you can create an array of arrays
	var board [8][8]string
	board[0][0] = "r"
	board[0][7] = "r"
	for column := range board[1] {
		board[1][column] = "p"
	}

	fmt.Println(board)
}

// This is a way to modify the values of an array,
// Keep in mind, a [5]string is not equal to [8]string despite of both being arrays and holding the same type
func modifyArray(array [9]string) [9]string {
	newArray := array
	for i := range newArray {
		newArray[i] = "New " + newArray[i]
	}
	return newArray
}
