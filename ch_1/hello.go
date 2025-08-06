package main

import "fmt"

func concat() {
	var username string = "John"
	var message string = "Hello, " + username + "!"
	fmt.Println(message)
}

func main() {
	fmt.Println("Hello, World!")
	concat()
}
