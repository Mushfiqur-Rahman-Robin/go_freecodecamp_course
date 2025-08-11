package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mux    sync.Mutex // no need for pointer to mutex
}

func (sc *safeCounter) inc(key string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.slowIncrement(key)
}

func (sc *safeCounter) val(key string) int {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	return sc.counts[key]
}

func (sc *safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

type emailTest struct {
	email string
	count int
}

// ===================== RWMutex Version =====================

type safeCounter2 struct {
	counts map[string]int
	mux    sync.RWMutex
}

func (sc *safeCounter2) inc(key string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.slowIncrement(key)
}

func (sc *safeCounter2) val(key string) int {
	sc.mux.RLock()
	defer sc.mux.RUnlock()
	return sc.counts[key]
}

func (sc *safeCounter2) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

type emailTest2 struct {
	email string
	count int
}

// ===================== Test Functions =====================

func test(sc *safeCounter, emailTests []emailTest) {
	emails := make(map[string]struct{})
	var wg sync.WaitGroup

	for _, emailT := range emailTests {
		emails[emailT.email] = struct{}{}
		for i := 0; i < emailT.count; i++ {
			wg.Add(1)
			go func(emailT emailTest) {
				sc.inc(emailT.email)
				wg.Done()
			}(emailT)
		}
	}
	wg.Wait()

	emailsSorted := make([]string, 0, len(emails))
	for email := range emails {
		emailsSorted = append(emailsSorted, email)
	}
	sort.Strings(emailsSorted)

	for _, email := range emailsSorted {
		fmt.Printf("Email: %s has %d emails\n", email, sc.val(email))
	}
	fmt.Println("=====================================")
}

func testRWMutex(sc *safeCounter2, emailTests []emailTest2) {
	emails := make(map[string]struct{})
	var wg sync.WaitGroup

	for _, emailT := range emailTests {
		emails[emailT.email] = struct{}{}
		for i := 0; i < emailT.count; i++ {
			wg.Add(1)
			go func(emailT emailTest2) {
				sc.inc(emailT.email)
				wg.Done()
			}(emailT)
		}
	}
	wg.Wait()

	emailsSorted := make([]string, 0, len(emails))
	for email := range emails {
		emailsSorted = append(emailsSorted, email)
	}
	sort.Strings(emailsSorted)

	for _, email := range emailsSorted {
		fmt.Printf("Email: %s has %d emails\n", email, sc.val(email))
	}
	fmt.Println("=====================================")
}

// ===================== Main =====================

func main() {
	sc := &safeCounter{
		counts: make(map[string]int),
	}
	test(sc, []emailTest{
		{"john@example.com", 23},
		{"john@example.com", 29},
		{"jill@example.com", 31},
		{"jill@example.com", 67},
	})
	test(sc, []emailTest{
		{"kaden@example.com", 23},
		{"george@example.com", 126},
		{"kaden@example.com", 31},
		{"george@example.com", 453},
	})

	sc2 := &safeCounter2{
		counts: make(map[string]int),
	}
	testRWMutex(sc2, []emailTest2{
		{"john@example.com", 23},
		{"john@example.com", 29},
		{"jill@example.com", 31},
		{"jill@example.com", 67},
	})
	testRWMutex(sc2, []emailTest2{
		{"kaden@example.com", 23},
		{"george@example.com", 126},
		{"kaden@example.com", 31},
		{"george@example.com", 453},
	})
}
