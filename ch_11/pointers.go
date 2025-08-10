package main

import (
	"fmt"
	"strings"
)

type Message struct {
	Recipient string
	Text      string
}

func sendMessage(m Message) {
	fmt.Printf("To: %v\n", m.Recipient)
	fmt.Printf("Message: %v\n", m.Text)
}

// pointer
func removeProfanity(message *string) {
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
	messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
	*message = messageVal
}

// pointer receivers
func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func test(recipient string, text string) {
	m := Message{Recipient: recipient, Text: text}
	sendMessage(m)
	fmt.Println("=====================================")
}

func testProfanity(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

func testPointerReceiver(e *email, newMessage string) {
	fmt.Println("-- before --")
	fmt.Println(e.message)
	fmt.Println("-- end before --")
	e.setMessage(newMessage)
	fmt.Println("-- after --")
	fmt.Println(e.message)
	fmt.Println("-- end after --")
	fmt.Println("==========================")
}

func main() {
	test("Lane", "Textio is getting better everyday!")
	test("Allan", "This pointer stuff is weird...")
	test("Tiffany", "What time will you be home for dinner?")

	messages1 := []string{
		"well shoot, this is awful",
		"dang robots",
		"dang them to heck",
	}

	messages2 := []string{
		"well shoot",
		"Allan is going straight to heck",
		"dang... that's a tough break",
	}

	testProfanity(messages1)
	testProfanity(messages2)

	testPointerReceiver(&email{
		message:     "this is my first draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my second draft")

	testPointerReceiver(&email{
		message:     "this is my third draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my fourth draft")
}
