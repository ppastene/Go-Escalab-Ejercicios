package main

import "fmt"

var a string

var (
	b int  = 44
	c bool = true
)

func main() {
	var age int = 30
	a = "This variable is a string"
	d := "Could be anything"
	fmt.Println(a, b, c, d, age)
	d = "Is another string!"
	fmt.Println(d)
	age++
	fmt.Println(age)
	age = age + 1
	fmt.Println(age)
	age += 1
	fmt.Println(age)

}
