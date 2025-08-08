package main

import "fmt"

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(msgtosend messageToSend) bool {
	if msgtosend.sender.name == "" {
		return false
	}
	if msgtosend.sender.number == 0 {
		return false
	}
	if msgtosend.recipient.name == "" {
		return false
	}
	if msgtosend.recipient.number == 0 {
		return false
	}
	return true
}

func test(mToSend messageToSend) {
	fmt.Printf(`sending "%s" from %s (%v) to %s (%v)...`,
		mToSend.message,
		mToSend.sender.name,
		mToSend.sender.number,
		mToSend.recipient.name,
		mToSend.recipient.number,
	)
	fmt.Println()
	if canSendMessage(mToSend) {
		fmt.Println("...sent!")
	} else {
		fmt.Println("...can't send message")
	}
	fmt.Println("====================================")
}

// embedded structs
type sender struct {
	rateLimit int
	newuser
}

type newuser struct {
	name   string
	number int
}

func testEmbedded(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("====================================")
}

func main() {
	test(messageToSend{
		message: "you have an appointment tommorow",
		sender: user{
			name:   "Brenda Halafax",
			number: 16545550987,
		},
		recipient: user{
			name:   "Sally Sue",
			number: 19035558973,
		},
	})
	test(messageToSend{
		message: "you have an event tommorow",
		sender: user{
			number: 16545550987,
		},
		recipient: user{
			name:   "Suzie Sall",
			number: 0,
		},
	})
	test(messageToSend{
		message: "you have an party tommorow",
		sender: user{
			name:   "Njorn Halafax",
			number: 16545550987,
		},
		recipient: user{
			name:   "Sally Sue",
			number: 19035558973,
		},
	})
	test(messageToSend{
		message: "you have a birthday tommorow",
		sender: user{
			name:   "Eli Halafax",
			number: 0,
		},
		recipient: user{
			name:   "Whitaker Sue",
			number: 19035558973,
		},
	})

	testEmbedded(sender{
		rateLimit: 50,
		newuser: newuser{
			name:   "John Doe",
			number: 16545550987,
		},
	})
	testEmbedded(sender{
		rateLimit: 0,
		newuser: newuser{
			name:   "John Doe",
			number: 16545550987,
		},
	})
}
