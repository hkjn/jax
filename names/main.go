package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//fmt.Printf("hello\n")
	//fmt.Printf("my name is Bobby\n")

	names := []string{
		"Anoushka",
		"Chuck",
		"Kimberly",
		"Ahmed",
	}
	now := int64(time.Now().Nanosecond())
	rand.Seed(now)
	printGreeting(names[rand.Intn(4)])
	printGreeting(names[rand.Intn(4)])
	printGreeting(names[rand.Intn(4)])
	printGreeting(names[rand.Intn(4)])
	//for i, name := range names {
	//	fmt.Printf("On item %d: %s\n", i, name)
	//	printGreeting(name)
	//}

}

func printGreeting(name string) {
	fmt.Printf("Hello!\n")
	fmt.Printf("My name  is %s\n", name)

}
