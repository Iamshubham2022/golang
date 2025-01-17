package main

import (
	"fmt"
	"sort"
	"encoding/json"
	"sync"
)

// Define a basic struct
type Student struct {
	Name  string
	Grade int
	Age   int
}

// Method for Student struct with value receiver
func (s Student) DisplayInfo() {
	fmt.Printf("Name: %s, Grade: %d, Age: %d\n", s.Name, s.Grade, s.Age)
}

// Method for Student struct with pointer receiver
func (s *Student) IncreaseGrade(percentage int) {
	s.Grade += percentage
}

// Struct with embedded fields
type Address struct {
	City    string
	Country string
}

type Person struct {
	Name    string
	Age     int
	Address // Embedding Address struct
}

// Using struct tags
type PersonWithTag struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

func main() {
	// 1. Basic struct creation and accessing fields
	student := Student{Name: "Alice", Grade: 90, Age: 20}
	student.DisplayInfo() // Output: Name: Alice, Grade: 90, Age: 20

	// Modify the struct fields
	student.Age = 21
	student.IncreaseGrade(10) // Increase grade by 10
	student.DisplayInfo() // Output: Name: Alice, Grade: 100, Age: 21

	// 2. Pointer to a struct and modifying via pointer
	studentPointer := &student
	studentPointer.Age = 22
	studentPointer.IncreaseGrade(5)
	student.DisplayInfo() // Output: Name: Alice, Grade: 105, Age: 22

	// 3. Array of structs
	students := []Student{
		{Name: "Bob", Grade: 85, Age: 22},
		{Name: "Charlie", Grade: 88, Age: 23},
	}
	for _, s := range students {
		s.DisplayInfo()
	}

	// 4. Slice of structs
	sliceOfStudents := []Student{
		{Name: "Dave", Grade: 91, Age: 20},
		{Name: "Eve", Grade: 76, Age: 21},
	}
	for _, s := range sliceOfStudents {
		s.DisplayInfo()
	}

	// 5. Map of structs
	studentMap := map[string]Student{
		"Alice":  {Name: "Alice", Grade: 100, Age: 22},
		"Bob":    {Name: "Bob", Grade: 85, Age: 22},
		"Charlie": {Name: "Charlie", Grade: 88, Age: 23},
	}
	for key, student := range studentMap {
		fmt.Printf("Key: %s -> ", key)
		student.DisplayInfo()
	}

	// 6. Using an anonymous struct
	anonymousStudent := struct {
		Name  string
		Grade int
	}{
		Name:  "John",
		Grade: 90,
	}
	fmt.Printf("Anonymous student: Name: %s, Grade: %d\n", anonymousStudent.Name, anonymousStudent.Grade)

	// 7. Nested structs and embedding
	person := Person{
		Name: "John",
		Age:  30,
		Address: Address{
			City:    "New York",
			Country: "USA",
		},
	}
	fmt.Printf("Person Name: %s, Age: %d, City: %s, Country: %s\n", person.Name, person.Age, person.City, person.Country)

	// 8. Struct tags and JSON encoding
	personWithTag := PersonWithTag{
		Name:    "Alice",
		Age:     25,
		Country: "USA",
	}
	jsonData, _ := json.Marshal(personWithTag)
	fmt.Println("PersonWithTag JSON:", string(jsonData)) // Output: {"name":"Alice","age":25,"country":"USA"}

	// 9. Methods on structs
	// Struct with method
	student2 := Student{Name: "Michael", Grade: 85, Age: 23}
	student2.DisplayInfo() // Output: Name: Michael, Grade: 85, Age: 23
	student2.IncreaseGrade(15)
	student2.DisplayInfo() // Output: Name: Michael, Grade: 100, Age: 23

	// 10. Sorting structs (by grade)
	studentsForSort := []Student{
		{Name: "Alice", Grade: 90, Age: 20},
		{Name: "Bob", Grade: 85, Age: 22},
		{Name: "Charlie", Grade: 88, Age: 23},
	}

	// Sorting by Grade
	sort.Slice(studentsForSort, func(i, j int) bool {
		return studentsForSort[i].Grade > studentsForSort[j].Grade
	})
	fmt.Println("Sorted students by grade:")
	for _, student := range studentsForSort {
		student.DisplayInfo()
	}

	// 11. Concurrency-safe map with sync.Map
	var sm sync.Map
	sm.Store("Alice", Student{Name: "Alice", Grade: 90, Age: 20})
	sm.Store("Bob", Student{Name: "Bob", Grade: 85, Age: 22})

	// Retrieving from sync.Map
	value, ok := sm.Load("Alice")
	if ok {
		student := value.(Student)
		fmt.Printf("Found in sync.Map: %s, Grade: %d, Age: %d\n", student.Name, student.Grade, student.Age)
	}

	// 12. Copying a map of structs
	originalMap := map[string]Student{
		"John": {Name: "John", Grade: 92, Age: 24},
		"Jane": {Name: "Jane", Grade: 87, Age: 21},
	}
	copyMap := make(map[string]Student)
	for k, v := range originalMap {
		copyMap[k] = v
	}

	// Print copied map
	fmt.Println("Copied Map:")
	for key, student := range copyMap {
		fmt.Printf("Key: %s -> Name: %s, Grade: %d, Age: %d\n", key, student.Name, student.Grade, student.Age)
	}
}
