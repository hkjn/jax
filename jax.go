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

func (jax Jax) say(something string) {
	fmt.Printf("%s says: %q\n", jax.name, something)
}

// feed gives Jax some food with a given value.
func (jax *Jax) feed(food string, value int) {
	jax.say("*chomp chomp*")
	jax.happiness += value
	fmt.Printf("Our jax ate the %q\n", food)
}

func (jax Jax) howAreYou() {
	msg := ""
	if jax.happiness > 100 {
		msg = fmt.Sprintf("Why I am just splendid!")
	} else {
		msg = fmt.Sprintf("I am fine, thank you.")
	}
	jax.say(msg)
}

func main() {
	// Declare a new variable 'j1' of type Jax, with 'happiness' set to 100 and 'name' set to "Jaxieboy".
	j1 := Jax{
		happiness: 30,
		name:      "Jaxieboy",
	}
	j1.say("Hello hello!")
	j1.say("Jax wants some food!")
	j1.feed("Bananas", 20)

	j2 := Jax{
		happiness: 90,
		name:      "CutieJax",
	}
	j2.say("Oh!")
	j2.say("A new friend")
	j2.feed("Chicken", 40)

	j1.howAreYou()
	j2.howAreYou()
}
