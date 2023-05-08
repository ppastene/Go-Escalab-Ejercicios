package main

import "fmt"

const DISTANCE = 236500000000000000000000          //Distance in Kilometers
const KILOMETERS_PER_LIGHTYEAR = 94610000000000000 //How many kilometers per light-years

func main() {
	fmt.Println("The andromeda galaxy its", DISTANCE/KILOMETERS_PER_LIGHTYEAR, "light-years away from earth")
}
