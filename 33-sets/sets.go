package main

import (
	"fmt"
	"sort"
)

func main() {
	// A set is an array which their elements are unique
	// In Golang the set type doens't, exist but we can modify an array to make their elements unique using maps
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool)

	for _, t := range temperatures {
		set[t] = true
	}
	if set[-28.0] {
		fmt.Println("Set member")
	} else {
		fmt.Println("Element doesn't not exist")
	}

	fmt.Println(set)

	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}
	sort.Float64s(unique)
	fmt.Println(unique)
}
