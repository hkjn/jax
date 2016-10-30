// Package jax is some Go ML experiments.
package main

import "fmt"

// say is a function that takes a single argument called 'x', which is a string.
func say(something string) {
	fmt.Printf("Jax says: %s\n", something)
}

// function add has two arguments which are both integers, and it adds those two integers and returns the value of that sum.
func add(a int, b int) int {
	return a + b
}

func main() {
	say("Hello hello!")
	say("Jax wants some food!")
	result := add(10, 4)
	arit := fmt.Sprintf("10 + 4 = %d\n", result)
	say(arit)
}
