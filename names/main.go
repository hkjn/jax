package main

import "fmt"

func main() {
	//fmt.Printf("hello\n")
	//fmt.Printf("my name is Bobby\n")
	names := []string{
		"Anoushka",
		"Chuck",
		"Kimberly",
		"Ahmed",
	}
	for i, name := range names {
		fmt.Printf("On item %d: %s\n", i, name)
		printGreeting(name)
	}

}

func printGreeting(name string) {
	fmt.Printf("Hello!\n")
	fmt.Printf("My name  is %s\n", name)

}
