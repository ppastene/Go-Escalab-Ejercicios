package main

import "fmt"

func main() {

	// An array and a slice have a sequential number as their index
	// An index in a map can be of any type, similar to hashes in Ruby, dictionaries in Python and objects in Javascript
	// A map must be declared like map[type_of_key]type_of_value{}
	temperatures := map[string]int{
		"earth": 15,
		"mars":  -65,
	}
	earth := temperatures["earth"]
	fmt.Printf("The average temperature on earth is %v°C\n", earth)
	// You can modify an element by calling its key
	temperatures["earth"] = 16
	// You can create a new element just by passing a new key and value
	temperatures["venus"] = 464
	fmt.Println(temperatures)
	// If you're calling an element that doesn't exist in the map it will show the default value
	// In the case of an integer it will be 0
	moon := temperatures["moon"]
	fmt.Println(moon)

	// You can use a 'comma ok' to see if an element of the map exists, if ok is true so the element exists
	if jupiter, ok := temperatures["jupiter"]; ok {
		fmt.Printf("The average temperature on jupiter is %v°C\n", jupiter)
	} else {
		fmt.Println("Where is jupiter?")
	}

	// A map doesnt copy when is assigned to a new variable,
	// and when it goes as an argument to a function or method the changes will be reflected outside
	planets := map[string]string{
		"earth": "Sector ZZ9",
		"moon":  "Sector ZZ9",
	}
	planetsMarkII := planets
	planets["earth"] = "Wrong Sector"
	fmt.Println(planets)
	fmt.Println(planetsMarkII)
	// Let's try with deleting the element
	delete(planets, "earth")
	fmt.Println(planetsMarkII)

	// To copy a map instead of using it as a reference we use the make() function
	// Accepts 2 arguments, the map and optionally the ammount of elements that are included
	// The initial value of it will be 0
	scores := make(map[string]int, 3)
	fmt.Println("Map:", scores, "LEN:", len(scores))
	scores["smitty"] = 100
	scores["paulie"] = 97
	scores["andrea"] = 92
	fmt.Println("Map:", scores, "LEN:", len(scores))

	// We can use maps to count anything
	times := []float64{
		10.0, 12.0, 9.0, 10.0, 13.0, 10.0, 15.0, 9.0,
	}
	frequency := make(map[float64]int)
	for _, t := range times {
		frequency[t]++
	}
	for t, num := range frequency {
		fmt.Printf("%.2f occurs %d times\n", t, num)
	}
	fmt.Println(frequency)

}
