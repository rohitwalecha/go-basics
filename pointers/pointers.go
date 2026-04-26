package main

import "fmt"

func main() {
	age := 32
	var agePointer *int
	agePointer = &age

	*agePointer-- // Used * to turn it back to its value

	fmt.Println(agePointer)  // The actual pointer
	fmt.Println(*agePointer) // Use * to get value of the pointer
	fmt.Println(age)         // The *agePointer and age should hold the same value
}
