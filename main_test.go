package main

import (
	"fmt"
)

func ExampleMain() {
	main()

	// Output:
	// hello world!
	// goodbye!
}

func ExampleGreeting() {
	fmt.Println(Greeting())

	// Output:
	// hello world!
}

func ExampleClosing() {
	fmt.Println(Closing())

	// Output:
	// goodbye!
}
