// In Go, the defer keyword is used to ensure that a function call is executed after
//  the surrounding function completes, but before the function returns. 
//  This can be useful for clean-up tasks such as closing files, releasing resources,
//   unlocking mutexes, or printing final messages.









// Detailed Explanation of Each Section:
// Basic Defer Example: Demonstrates the use of defer to execute a function after the main function execution.

// defer fmt.Println("This is deferred") will be printed after the main logic.
// Multiple Defer Statements: Shows that multiple defer calls execute in reverse order (LIFO).

// The deferred functions are executed after the normal execution of the main() function finishes.
// Defer with Return Statement: Demonstrates how a deferred function is executed before the return statement in the function is processed.

// It prints "Deferred function executed" before the function returns a value.
// Defer with Error Handling (Resource Management): Used to ensure that resources like files are properly closed, even if there is an error.

// The file is opened, and defer file.Close() ensures that the file is closed at the end of the function.
// Defer with Mutex (Concurrency Example): Shows how defer is useful when working with mutexes for concurrent programming.

// The mutex is locked at the beginning and unlocked at the end of the function.
// Measuring Function Execution Time using Defer: Uses defer to measure the time it takes for a function to execute.

// The function funcWithTimeMeasurement uses defer to log the time after it finishes processing.
// Defer in Loops: Demonstrates that when defer is used inside a loop, the deferred function calls are executed in reverse order.

// It prints numbers 2, 1, 0, after the loop completes.
// Defer and Resource Cleanup: A generic cleanup example where resources like memory or network connections can be cleaned up using defer.

// Defer with Panic and Recover (Error Handling): This demonstrates how defer can be used with recover to handle panics and prevent the program from crashing.

// The panic is triggered, but recover catches it and prints the error message.
// Key Concepts:
// LIFO Execution: defer calls are executed in Last In, First Out order.
// Resource Management: defer is often used to clean up resources like files or network connections (i.e., closing files).
// Mutex Unlocking: In concurrent programming, defer ensures that mutexes are unlocked after being locked, preventing deadlocks.
// Error Handling: Using defer with recover to handle unexpected panics.
// Timing Execution: It can be used to measure the execution time of a function by recording the start time and printing the duration when the function exits.
// Important Considerations:
// Performance: While defer is a great tool for clean-up and managing resources, it has a small performance overhead. For time-critical operations in highly-concurrent code, defer should be used carefully.
// Multiple Defers: If you use multiple defer statements, they will execute in reverse order (LIFO), which can be handy in some scenarios, like when locking and unlocking resources.
// Let me know if you need any further details or examples!


package main

import (
	"fmt"
	"os"
	"time"
	"sync"
)

func main() {
	// 1. Basic Defer Example
	fmt.Println("1. Basic Defer Example:")
	defer fmt.Println("This is deferred")
	fmt.Println("End of main function")
	
	// 2. Multiple Defer Statements
	fmt.Println("\n2. Multiple Defer Statements:")
	defer fmt.Println("First deferred function")
	defer fmt.Println("Second deferred function")
	defer fmt.Println("Third deferred function")
	fmt.Println("Main function execution")
	
	// 3. Defer with Return Statement
	fmt.Println("\n3. Defer with Return Statement:")
	funcWithDefer()

	// 4. Defer with Error Handling (Resource Management)
	fmt.Println("\n4. Defer with Error Handling (Resource Management):")
	file, err := os.Open("example.txt") // This file must exist for this to work
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	fmt.Println("File opened successfully")
	
	// 5. Defer with Mutex (Concurrency Example)
	fmt.Println("\n5. Defer with Mutex (Concurrency Example):")
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("Mutex locked and unlocked")
	
	// 6. Measuring Function Execution Time using Defer
	fmt.Println("\n6. Measuring Function Execution Time using Defer:")
	funcWithTimeMeasurement()

	// 7. Defer in Loops
	fmt.Println("\n7. Defer in Loops:")
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	// 8. Defer and Resource Cleanup
	fmt.Println("\n8. Defer and Resource Cleanup:")
	defer resourceCleanup()
	fmt.Println("Resource cleanup will be executed when function exits")

	// 9. Defer with Panic and Recover (Error Handling)
	fmt.Println("\n9. Defer with Panic and Recover:")
	funcWithPanicAndRecover()
}

// 3. Defer with Return Statement
func funcWithDefer() int {
	defer fmt.Println("Deferred function executed")
	return 42
}

// 6. Measuring Function Execution Time using Defer
func funcWithTimeMeasurement() {
	start := time.Now()
	defer func() {
		fmt.Println("Function executed in", time.Since(start))
	}()
	// Simulate some work with a sleep
	time.Sleep(2 * time.Second)
}

// 8. Defer and Resource Cleanup
func resourceCleanup() {
	fmt.Println("Cleaning up resources...")
}

// 9. Defer with Panic and Recover (Error Handling)
func funcWithPanicAndRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("Function started")
	panic("Something went wrong!")
}
