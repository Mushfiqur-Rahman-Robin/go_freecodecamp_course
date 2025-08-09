package main

import (
	"fmt"
)

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1 + (float64(i) * 0.01)
	}
	return totalCost
}

// returns the maximum number of messages that can be sent
func maxMessages(thresh float64) int {
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost += 1 + (float64(i) * 0.01)
		if totalCost > thresh {
			return i
		}
	}
}

// while like loop (go does not have explicit while loops)
func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 0
	for actualCostInPennies <= float64(maxCostInPennies) {
		maxMessagesToSend++
		actualCostInPennies *= costMultiplier
	}
	return maxMessagesToSend
}

// fizzbuzz
func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// continue and break
func printPrimes(max int) {
	for n := 2; n <= max; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}

		if n%2 == 0 {
			continue
		}

		isPrime := true
		for i := 3; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(n)
		}
	}
}

func test(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
	fmt.Println("===============================================================")
}

func test2(thresh float64) {
	fmt.Printf("Threshold: %.2f\n", thresh)
	max := maxMessages(thresh)
	fmt.Printf("Maximum messages that can be sent: = %v\n", max)
	fmt.Println("===============================================================")
}

func test3(costMultiplier float64, maxCostInPennies int) {
	maxMessagesToSend := getMaxMessagesToSend(costMultiplier, maxCostInPennies)
	fmt.Printf("Multiplier: %v\n",
		costMultiplier,
	)
	fmt.Printf("Max cost: %v\n",
		maxCostInPennies,
	)
	fmt.Printf("Max messages you can send: %v\n",
		maxMessagesToSend,
	)
	fmt.Println("====================================")
}

func testPrimes(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
	test(40)
	test(50)

	test2(10.0)
	test2(20.0)
	test2(30.0)
	test2(40.0)
	test2(50.0)

	test3(1.1, 5)
	test3(1.2, 10)
	test3(1.3, 100)

	fizzbuzz()
	testPrimes(100)
}
