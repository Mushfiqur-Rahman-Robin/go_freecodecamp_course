package main

import (
	"errors"
	"fmt"
)

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

func yearsUntilEvents(age int) (
	yearsUntilAdult int, yearsUntilDriving int, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age

	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}

	yearsUntilDriving = 16 - age
	if yearsUntilDriving < 0 {
		yearsUntilDriving = 0
	}

	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}

	return yearsUntilAdult, yearsUntilDriving, yearsUntilCarRental
}

func test(age int) {
	fmt.Println("Age: ", age)
	yearsUntilAdult, yearsUntilDriving, yearsUntilCarRental := yearsUntilEvents(age)

	fmt.Println("You will be an adult in", yearsUntilAdult, "years")
	fmt.Println("You will be able to drive in", yearsUntilDriving, "years")
	fmt.Println("You will be able to rent a car in", yearsUntilCarRental, "years")
}

func divide(divident int, divisor int) (int, error) { // guard clause
	if divisor == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return divident / divisor, nil
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

	test(21)
	test(16)

	divide(10, 2)
	divide(10, 0)
}
