package main

import (
	"fmt"
)

func main() {
	fmt.Println("Drawing of a chessboard with their pieces using arrays, runes and loops")
	fmt.Println("Only for the intention to show how to iterate an array, maybe in a future I can add some logic")
	fmt.Println("Uppercase: White - Lowercase: Black - #: Black square - p: pawn - r: roke - c: knight - b: bishop - q: queen - k: king")

	var chessPieces [8][8]rune
	chessPieces[0] = [8]rune{
		114, 99, 98, 113, 107, 98, 99, 114,
	}
	chessPieces[7] = [8]rune{
		82, 67, 66, 81, 75, 66, 67, 82,
	}
	for column := range chessPieces[1] {
		chessPieces[1][column] = 112
	}
	for column := range chessPieces[6] {
		chessPieces[6][column] = 80
	}

	drawBoard(chessPieces)
}

func drawBoard(pieces [8][8]rune) {
	fmt.Println("   _______________")
	for i := 0; i < 8; i++ {
		fmt.Printf("%v |", 8-i)
		for j := 0; j < 8; j++ {
			switch pieces[i][j] {
			case 0:
				if ((j+1)%2 == 0 && (i+1)%2 == 1) || ((j+1)%2 == 1 && (i+1)%2 == 0) {
					fmt.Print("#|")
				} else {
					fmt.Print("_|")
				}

			default:
				fmt.Printf("%c|", pieces[i][j])
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println("   a b c d e f g h ")
}
