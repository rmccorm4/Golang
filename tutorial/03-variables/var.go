package main

import "fmt"

func main() {
	// Variable declaration with type at the end
	var age int
	// Show that go sets default value to the "zero" of its type
	fmt.Println("My default age is", age)

	// Assign value
	age = 20
	fmt.Println("My current age is", age)

	// Doing it all at once
	var curr_age int = 20
	fmt.Println("My, efficiently assigned, current age is", curr_age)

	// Types can also be inferred in Go based on assignment
	var inferred_age = 20
	fmt.Println("My -type inferred- age is", inferred_age)

	// Variables can also be assigned in the same line
	var height, width int = 4, 2
	fmt.Println("Height =", height, "Width =", width)

	// Types can still be inferred with multiple assignment
	var h, w = 10, 5
	fmt.Println("New Height =", h, "New Width =", w)
	
	// Let's see if multiple assignment type infernce works for
	// variables of different types
	var name, rank = 'Ryan', 1
	fmt.Println("Your name is", name, "and you are number ", rank, "!")
}
