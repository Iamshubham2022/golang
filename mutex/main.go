// Mutex (Mutual Exclusion) in Golang
// Mutex multiple goroutines ke beech race condition avoid karne ke liye use hota hai. 
// Ye ek lock mechanism provide karta hai jo ensure karta hai 
// ki ek time par sirf ek goroutine shared resource ko access kare.

// go run -- race condition

package main
// package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
	mutex.Lock()   // Lock kar diya
	counter++      // Critical section
	mutex.Unlock() // Unlock kar diya
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter) // Expected Output: 10
}
