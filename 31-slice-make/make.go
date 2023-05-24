package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v \n", label, len(slice), cap(slice), slice)
}

func main() {
	// make() allows to set manually the length and capacity of an array
	// The argument after the array is the length
	names := make([]string, 10)
	dump("names", names)
	// The next one is the capacity
	// Keep in mind, after the array reaches its capacity it will duplicate by its value
	countries := make([]string, 10, 20)
	countries = append(countries, "Rand McNally")
	dump("countries", countries)

	planets := terraform("New", "Venus", "Earth", "Mars")
	fmt.Println(planets)

	dwarfs := []string{"Ceres", "Pluto", "Haumea"}
	newDwarfs := terraform("Amazing", dwarfs...)
	fmt.Println(dwarfs)
	fmt.Println(newDwarfs)
}

// This function returns a new slice with its array
// With the ... is possible to insert many arguments as possible so its possible to get the length of arguments
// and set it to the array
func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}

	return newWorlds
}
