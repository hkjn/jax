// This is Aruna learning how to code in go
package main

//Allows me to use the fmt pakage which can format things and print them to the terminal
import "fmt"

func main() {
	//We will run this function when we start the program
	fmt.Println("oivey")
	q := f(8)
	fmt.Printf("fucking q is %d!?!?!?\n", q)
}

func f(x int) int {
	y := x * x
	return y
}
