package main

import "fmt"

func main() {
	// 1. Declaring and Initializing a Pointer
	var ptr *int
	fmt.Println("Pointer value (uninitialized):", ptr) // nil, as it is not pointing to anything

	// Initializing a variable and pointing to it
	num := 42
	ptr = &num // ptr now holds the address of num
	fmt.Println("Pointer value (address):", ptr)
	fmt.Println("Value stored at the address:", *ptr) // Dereferencing to get the value

	// 2. Modifying Value through a Pointer
	*ptr = 100 // Changing the value at the memory location pointed by ptr
	fmt.Println("Modified num through pointer:", num)

	// 3. Pointer to Pointer
	var ptrToPtr **int
	ptrToPtr = &ptr // ptrToPtr points to ptr
	fmt.Println("Pointer to Pointer:", ptrToPtr)
	fmt.Println("Value via Pointer to Pointer:", **ptrToPtr)

	// 4. Passing Pointers to Functions
	increment(&num)
	fmt.Println("After increment function, num:", num)

	// 5. Returning Pointers from Functions
	newPtr := createPointer(55)
	fmt.Println("Value from returned pointer:", *newPtr)

	// 6. Pointer and Slices
	slice := []int{1, 2, 3}
	modifySlice(slice) // Slices are reference types; changes reflect in the original slice
	fmt.Println("Modified slice:", slice)

	// 7. Pointer to Struct
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "Alice", Age: 30}
	personPtr := &person
	fmt.Println("Original person:", *personPtr)
	personPtr.Age = 31 // Modifying struct fields via pointer
	fmt.Println("Modified person:", *personPtr)

	// 8. Pointer Arithmetic (Not allowed in Go)
	// Unlike languages like C or C++, Go does not support pointer arithmetic for safety reasons.

	// 9. Nil Pointer
	var nilPtr *int
	if nilPtr == nil {
		fmt.Println("nilPtr is nil")
	}

	// 10. Array of Pointers
	arr := [3]int{10, 20, 30}
	var arrPointers [3]*int
	for i := 0; i < len(arr); i++ {
		arrPointers[i] = &arr[i]
	}
	fmt.Println("Array of Pointers:")
	for i, ptr := range arrPointers {
		fmt.Printf("Index %d, Address: %p, Value: %d\n", i, ptr, *ptr)
	}

	// 11. Pointer and Maps (Maps are reference types; no direct use of pointers)
	m := map[string]int{"a": 1, "b": 2}
	modifyMap(m)
	fmt.Println("Modified map:", m)

	// 12. Pointer and Concurrency (Syncing with goroutines)
	num = 200
	go modifyValue(&num) // Modifies value concurrently
	fmt.Println("Value after concurrent modification (might vary):", num) // Unreliable without sync
}

// Function to increment a value using pointer
func increment(n *int) {
	*n++
}

// Function to create and return a pointer
func createPointer(value int) *int {
	ptr := new(int) // Allocates memory for an int
	*ptr = value    // Sets the value
	return ptr
}

// Function to modify a slice
func modifySlice(s []int) {
	for i := range s {
		s[i] *= 2
	}
}

// Function to modify a map
func modifyMap(m map[string]int) {
	m["a"] = 10
	m["c"] = 3 // Adding a new key-value pair
}

// Function to modify a value concurrently
func modifyValue(n *int) {
	*n += 50
}
