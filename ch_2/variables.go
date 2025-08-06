package main

import "fmt"

func main() {
	var smsSendingLimit int // this is variable declaration but no initialization
	var costPerMessage float64

	smsSendingLimit = 50
	costPerMessage = 0.0075
	wishSentence := "Have a nice day!" // this is variable declaration

	const name = "John"
	const age = 30
	const pi float64 = 3.14159

	minMessageLength := 10
	maxMessageLength := 20

	res := fmt.Sprintf("My name is %s and I'm %d years old", name, age)

	fmt.Println(res)
	fmt.Println("The value of pi is: ", pi)

	fmt.Println("SMS sending limit: ", smsSendingLimit)
	fmt.Println("Cost per message: ", costPerMessage)
	fmt.Println(wishSentence)
	fmt.Println("The type of wishSentence is: ", fmt.Sprintf("%T", wishSentence))

	if minMessageLength < maxMessageLength {
		fmt.Println("Message length is between", minMessageLength, "and", maxMessageLength)
	}
}
