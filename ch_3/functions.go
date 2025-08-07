package main

import "fmt"

func concat(a string, b string) string {
	return a + b
}

func incrementSend(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}

func getNames() (string, string) {
	return "John", "Doe"
}

func main() {

	fmt.Println(concat("Hello, ", "World!"))
	fmt.Println(concat("Go is ", "awesome!"))

	sendsSoFar := 400
	const sendsToAdd = 50

	sendsSoFar = incrementSend(sendsSoFar, sendsToAdd)
	fmt.Println("You have sent", sendsSoFar, "messages")

	firstName, _ := getNames() // _ is a blank identifier, it is used to ignore the second return value as go does not allow unused variables
	fmt.Println("Your name is", firstName)
}
