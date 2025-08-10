package main

import (
	"errors"
	"fmt"
	"sort"
)

func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}
	return formattedMessages
}

func addSignature(message string) string {
	return message + " Kind regards."
}

func addGreeting(message string) string {
	return "Hello! " + message
}

// currying
// getLogger takes a function that formats two strings into
// a single string and returns a function that formats two strings but prints
// the result instead of returning it
func getLogger(formatter func(string, string) string) func(string, string) {
	return func(first, second string) {
		fmt.Println(formatter(first, second))
	}
}

// defer
const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

func logAndDelete(users map[string]user, name string) (log string) {
	defer delete(users, name)

	user, ok := users[name]
	if !ok {
		return logNotFound
	}
	if user.admin {
		return logAdmin
	}
	return logDeleted
}

type user struct {
	name   string
	number int
	admin  bool
}

// closure
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

type emailBill struct {
	costInPennies int
}

func test(messages []string, formatter func(string) string) {
	defer fmt.Println("====================================")
	formattedMessages := getFormattedMessages(messages, formatter)
	if len(formattedMessages) != len(messages) {
		fmt.Println("The number of messages returned is incorrect.")
		return
	}
	for i, message := range messages {
		formatted := formattedMessages[i]
		fmt.Printf(" * %s -> %s\n", message, formatted)
	}
}

func testCurrying(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("====================================")
	logger := getLogger(formatter)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func colonDelimit(first, second string) string {
	return first + ": " + second
}
func commaDelimit(first, second string) string {
	return first + ", " + second
}

func testDefer(users map[string]user, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("====================================")
	log := logAndDelete(users, name)
	fmt.Println("Log:", log)
}

func testClosure(bills []emailBill) {
	defer fmt.Println("====================================")
	countAdder, costAdder := adder(), adder()
	for _, bill := range bills {
		fmt.Printf("You've sent %d emails and it has cost you %d cents\n", countAdder(1), costAdder(bill.costInPennies))
	}
}

func main() {
	test([]string{
		"Thanks for getting back to me.",
		"Great to see you again.",
		"I would love to hang out this weekend.",
		"Got any hot stock tips?",
	}, addSignature)
	test([]string{
		"Thanks for getting back to me.",
		"Great to see you again.",
		"I would love to hang out this weekend.",
		"Got any hot stock tips?",
	}, addGreeting)

	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}
	testCurrying("Error on database server", dbErrors, colonDelimit)

	mailErrors := []error{
		errors.New("email too large"),
		errors.New("non alphanumeric symbols found"),
	}
	testCurrying("Error on mail server", mailErrors, commaDelimit)

	// defer
	users := map[string]user{
		"john": {
			name:   "john",
			number: 18965554631,
			admin:  true,
		},
		"elon": {
			name:   "elon",
			number: 19875556452,
			admin:  true,
		},
		"breanna": {
			name:   "breanna",
			number: 98575554231,
			admin:  false,
		},
		"kade": {
			name:   "kade",
			number: 10765557221,
			admin:  false,
		},
	}

	fmt.Println("Initial users:")
	usersSorted := []string{}
	for name := range users {
		usersSorted = append(usersSorted, name)
	}
	sort.Strings(usersSorted)
	for _, name := range usersSorted {
		fmt.Println(" -", name)
	}
	fmt.Println("====================================")

	testDefer(users, "john")
	testDefer(users, "santa")
	testDefer(users, "kade")

	fmt.Println("Final users:")
	usersSorted = []string{}
	for name := range users {
		usersSorted = append(usersSorted, name)
	}
	sort.Strings(usersSorted)
	for _, name := range usersSorted {
		fmt.Println(" -", name)
	}
	fmt.Println("====================================")

	testClosure([]emailBill{
		{45},
		{32},
		{43},
		{12},
		{34},
		{54},
	})

	testClosure([]emailBill{
		{12},
		{12},
		{976},
		{12},
		{543},
	})

	testClosure([]emailBill{
		{743},
		{13},
		{8},
	})
}
