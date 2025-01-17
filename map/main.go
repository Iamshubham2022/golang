package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	// Create a map using make
	student := make(map[string]int)
	student["jai"] = 23
	student["kul"] = 45
	student["shub"] = 23
	student["rohi"] = 32

	fmt.Println("The grade of jai is:", student["jai"])

	// Iterate over the map
	for key, value := range student {
		fmt.Printf("The grade of %s is %d\n", key, value) 
	}

	// Check if a key exists in the map
	grade, exists := student["jai"]
	Grade, Exit := student["shu"]

	if exists {
		fmt.Println("The grade of jai:", grade)
	} else {
		fmt.Println("jai does not exist")
	}

	if Exit {
		fmt.Println("The grade of shu:", Grade)
	} else {
		fmt.Println("shu does not exist")
	}

	// Modify the map
	student["jai"] = 100
	fmt.Println("The updated grade of jai is:", student["jai"])

	// Delete a key from the map
	delete(student, "jai")
	_, exists = student["jai"]
	if exists {
		fmt.Println("The grade of jai is still present")
	} else {
		fmt.Println("The grade of jai has been deleted")
	}

	// Another map example with strings
	person := map[string]string{
		"name":    "shubh",
		"student": "yes",
		"rollno":  "21010",
	}
	for key, value := range person {
		fmt.Printf("The person %s is %s\n", key, value)
	}

	// Check if a key exists in another map
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	_, exists = m["key1"]
	if exists {
		fmt.Println("Key exists")
	} else {
		fmt.Println("Key does not exist")
	}

	// Copy a map
	original := map[string]int{"a": 1, "b": 2}
	copyMap := make(map[string]int)
	for k, v := range original {
		copyMap[k] = v
	}
	fmt.Println("Original map:", original)
	fmt.Println("Copied map:", copyMap)

	// Clear a map
	m = make(map[string]int) // Reinitialize
	fmt.Println("Cleared map:", m)

	// Structs in maps
	type Student struct {
		Name  string
		Grade int
	}

	students := map[string]Student{
		"1": {"Alice", 90},
		"2": {"Bob", 85},
	}

	for id, student := range students {
		fmt.Printf("Student ID: %s, Name: %s, Grade: %d\n", id, student.Name, student.Grade)
	}

	// Using structs as keys
	type Point struct {
		X, Y int
	}

	points := map[Point]string{
		{X: 1, Y: 2}: "Point A",
		{X: 3, Y: 4}: "Point B",
	}
	for point, name := range points {
		fmt.Printf("Point (%d, %d): %s\n", point.X, point.Y, name)
	}

	// Map of slices
	groupedNumbers := map[string][]int{
		"group1": {1, 2, 3},
		"group2": {4, 5, 6},
	}
	fmt.Println("Group1 numbers:", groupedNumbers["group1"])

	// Nested maps
	nestedMap := map[string]map[string]int{
		"outer": {"inner": 42},
	}
	fmt.Println("Nested map value:", nestedMap["outer"]["inner"])

	// Sorting map keys
	m = map[string]int{"a": 3, "c": 1, "b": 2}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("Key: %s, Value: %d\n", k, m[k])
	}

	// Concurrency-safe map using sync.Map
	var sm sync.Map
	sm.Store("key1", "value1")
	value, _ := sm.Load("key1")
	fmt.Println("Value from sync.Map:", value)
}
