// Cesar encoding: sumar 3
// ROT13: Sumar 13

package main

import (
	"fmt"
)

func main() {
	var input, caesar, rot13 string
	input = "This is a text to cipher into caesar and rot13 algorithms"
	for i := range input {
		c := input[i]
		r := input[i]
		if input[i] >= 'A' && input[i] <= 'z' {
			c += 3
			r += 13
			if c > 'z' {
				c = c - 26
			}
			if r > 'z' {
				r = r - 26
			}
		}

		caesar += string(c)
		rot13 += string(r)
	}

	fmt.Printf("Carsar: %v \n", caesar)
	fmt.Printf("ROT13: %v", rot13)
}
