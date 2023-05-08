package main

import (
	"fmt"
	"time"
)

func main() {
	month := time.Now().Month()
	today := time.Now().Day()
	fmt.Printf("The current month is %v\n", month.String())
	switch month {
	case 2:
		fmt.Println("February is the only month to have 28 days")
	case 4, 6, 9, 11:
		fmt.Println("April, June, September and November are the months to have 30 days")
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("January, March, May, July, August, October and December are the months to have 31 days")
	}
	fmt.Println("As you can see a switch in Go doesn't need the break instruction to exit the switch because does it automatically")
	fmt.Println("But, with the use of fallthrough we can access to the next case")
	fmt.Printf("Today is %d. There's something to do today?\n", today)
	switch today {
	case 5:
		fmt.Println("Clean your house.")
		fallthrough
	case 10:
		fmt.Println("Buy some wine.")
		fallthrough
	case 15:
		fmt.Println("Visit a doctor.")
		fallthrough
	case 25:
		fmt.Println("Buy some food.")
		fallthrough
	case 30, 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("Nothing to do today but rebember, be a good boy.")
	}
}
