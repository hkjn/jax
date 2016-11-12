// This is Aruna learning how to code in go
package sub

// Allows me to use the fmt pakage which can format things and print them to the terminal
import "fmt"

// main.main() ->
//   fmt.Println()
//   Square()
//   fmt.Printf()
//   Square()
//     Square()
//   fmt.Println()
//   fmt.Printf()
func main() {
	// We will run this function when we start the program
	fmt.Println("oivey")
	q := Square(8)
	fmt.Printf("fucking q is %d!?!?!?\n", q)
	k := Square(Square(100))
	fmt.Println("oibayayushki!!")
	fmt.Printf("k is %d!\n", k)
}

// Square returns the square of its argument
func Square(x int) int {
	return x * x
}
