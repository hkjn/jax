// Package jax is some Go ML experiments.
package main

import "fmt"

// Jax is a new type, which is a struct containing two fields:
// - an int called 'happiness'
// - a string called 'name'
type Jax struct {
	happiness int
	name      string
}

// say is a function that takes two arguments:
//   - one called 'jax' of type Jax
//   - one called 'something' of type string
func say(jax Jax, something string) {
	fmt.Printf("%s says: %q\n", jax.name, something)
	fmt.Printf("(I am this happy: %d)\n\n", jax.happiness)
}

func main() {
	// Declare a new variable 'j1' of type Jax, with 'happiness' set to 100 and 'name' set to "Jaxieboy".
	j1 := Jax{
		happiness: 100,
		name:      "Jaxieboy",
	}
	say(j1, "Hello hello!")
	say(j1, "Jax wants some food!")
	j1.happiness += 10
	say(j1, "*chomp chomp*")

	j2 := Jax{
		happiness: 90,
		name:      "CutieJax",
	}
	say(j2, "Oh!")
	say(j2, "A new friend")
}
