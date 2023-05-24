package main

import "fmt"

func main() {
	textToDecrypt := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	keyIndex := 0
	message := ""
	for _, letter := range textToDecrypt {
		k := rune(keyword[keyIndex]) - 'A'
		newLetter := letter - k
		if newLetter < 'A' {
			newLetter += 26
		}
		message += string(newLetter)
		keyIndex++
		keyIndex %= len(keyword)
	}

	fmt.Println(message)
}
