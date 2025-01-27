// The Wait() method blocks the execution of the main program until all the
//  goroutines associated with a sync.WaitGroup complete their tasks.
// It's typically used in conjunction with Add() (to add goroutines) and Done()
//  (to signal a goroutine is finished).







package main

import (
	"fmt"
	"sync"
)

func printsomthing(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	// time.Sleep(1000 * time.Millisecond)
	fmt.Printf("work start: %d\n", i)
	fmt.Printf("work end: %d\n", i)
}

func main() {
	fmt.Println("Start learning of sync functions and wait")
	var wg sync.WaitGroup
	for i := 1; i < 6; i++ {
		wg.Add(1)
		go printsomthing(i, &wg)
	}
	wg.Wait()

	fmt.Println("End learning of sync functions and wait")
}
