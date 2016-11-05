// This is Aruna learning how to code in go
package main

// Allows me to use the fmt pakage which can format things and print them to the terminal
import "fmt"

// main.main() ->
//   fmt.Println()
//   f()
//   fmt.Printf()
//   f()
//   fmt.Println()
//   fmt.Printf()
func main() {
	// We will run this function when we start the program
	fmt.Println("oivey")
	q := square(8)
	fmt.Printf("fucking q is %d!?!?!?\n", q)
	z := square(100)
	k := square(z)
	fmt.Println("oibayayushki!!")
	fmt.Printf("k is %d!\n", k)
}

func square(x int) int {
	y := x * x
	return y
}
