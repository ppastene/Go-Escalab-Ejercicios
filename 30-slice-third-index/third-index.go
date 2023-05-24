package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v \n", label, len(slice), cap(slice), slice)
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
	dump("planets", planets)
	// The third index limits the elements of the slice
	// Here the slice have 4 elements counting from index 0 and limits it to 4
	terrestrial := planets[0:4:4]
	dump("terrestrial", terrestrial)
	dump("planets without 3rd index", planets[0:4])
	// Here the slice got a new element with append() and now shows 5 elements
	worlds := append(terrestrial, "Ceres")
	dump("worlds", worlds)
	// Keep an eye, because a slice is pointing to the array can happend a data overwrite
	// Here we append the Ceres element into a planet[0:4] slice, so its become the 5th element
	// but since is still pointing to the planets array the value also changes the 5th element of the array
	dump("Append Ceres into Slice", append(planets[0:4], "Ceres"))
	dump("planets", planets)
}
