// Cesar encoding: sumar 3
// ROT13: Sumar 13

package main

import (
	"fmt"
	"strings"
)

func main() {
	var input, caesar, rot13 string
	input = "Cipher me please"
	input = strings.ToUpper(input)
	for i := range input {
		c := input[i]
		r := input[i]
		if input[i] >= 'A' && input[i] <= 'Z' {
			c += 3
			r += 13
			if c > 'Z' {
				c -= 26
			}
			if r > 'Z' {
				r -= 26
			}
		}

		caesar += string(c)
		rot13 += string(r)
	}

	fmt.Printf("Carsar: %v \n", caesar)
	fmt.Printf("ROT13: %v", rot13)
}
