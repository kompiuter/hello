package main

import "fmt"

func main() {
	fmt.Println(Greeting())
	fmt.Println(Closing())
}

func Greeting() string {
	return "hello world!"
}

func Closing() string {
	return "goodbye!"
}
