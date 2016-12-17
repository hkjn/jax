// Package jax is some Go ML experiments.
package jaxhelper

import (
	"fmt"
	"io"

	"hkjn.me/jax/sub"
)

var Chicken = Food{
	Name:  "Chicken",
	Value: 40,
}

var Banana = Food{
	Name:  "Banana",
	Value: 3,
}

// Jax is a new type, which is a struct containing two fields:
// - an int called 'happiness'
// - a string called 'name'
type Jax struct {
	Happiness int
	Name      string
	Writer    io.Writer
}

type Food struct {
	Name  string
	Value int
}

func (jax Jax) Say(something string) {
	x := sub.Square(jax.Happiness)
	fmt.Fprintf(jax.Writer, "%s says: %q (my happiness squared is %v)\n", jax.Name, something, x)
}

// feed gives Jax some food with a given value.
func (jax *Jax) Feed(food Food) {
	jax.Say("*chomp chomp*")
	jax.Happiness += food.Value
	fmt.Fprintf(jax.Writer, "%v ate the %q\n", jax.Name, food.Name)
}

func (jax Jax) HowAreYou() {
	msg := ""
	if jax.Happiness > 100 {
		msg = fmt.Sprintf("Why I am just splendid!")
	} else {
		msg = fmt.Sprintf("I am fine, thank you.")
	}
	jax.Say(msg)
}
