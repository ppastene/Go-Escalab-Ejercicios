package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v \n", label, len(slice), cap(slice), slice)
}

func main() {
	// The numbers of element of a slice shows its length len(), but also a slice can increase its capacity cap()
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	// dwarfs shows a length of 5 and a capacity of 5, which means it have 5 elements and a capacity to fill 5 elements
	// which are filled right now
	// dwarfs[1:2] shows a length of 1 element and a capacity of 4 despite of showing only one element,
	// that's because it start from the index 1 and shows the 2nd element which is indeed the index 1 so the length is 1
	// but since it start from the index 1 it counts the rest of elements of the array which are 4 elements including the index itself
	// so the capacity  of the slice is 4
	dump("dwarfs", dwarfs)
	dump("dwarfs[1:2]", dwarfs[1:2])

	// dwarfs1 shows a length of 5 and a capacity of 5,
	// in dwarfs2 it appends 1 element so it have a length of 6 elements and a capacity of 10.
	// What happend is the compiler creates a new array with the double of capacity to fill the previous array and the new element
	// So in dwarfs3 it appends 3 elements and now have a length of 9 elements and a capacity of 10.
	// If we add one more element we fill the capacity of the array,
	// but if we add another one it creates another array with the double of capacity, in this case 20
	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dump("dwarfs1", dwarfs1)
	dwarfs2 := append(dwarfs1, "Orcus")
	dump("dwarfs2", dwarfs2)
	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")
	dump("dwarfs3", dwarfs3)

	// The value of dwarfs3[1] changes also in dwarfs2 but no in dwarfs1
	// that's because for dwarfs2 Golang creates a new array from dwarfs1 to increase its capacity
	// and dwarfs3 uses the same array as dwarfs2.
	// So its ok to say the array of dwarfs1 is different of the array from dwarfs2 and dwarfs3
	dwarfs3[1] = "New Pluto"
	fmt.Println(dwarfs1)
	fmt.Println(dwarfs2)
	fmt.Println(dwarfs3)
}
