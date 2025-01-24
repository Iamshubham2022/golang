package main

import (
    "fmt"
    "reflect"
)

//In go lang ther is no while loop and do while loop, only for loop,
//you can perform all the opration of while loop in for loop,
// you can also use like break , continue etc..

func main() {
	// // how we wiil print an array by using for loop
	// ans := []int{13, 34, 53, 6, 73, 68, 9}

	// for index, value := range ans {
	// 	fmt.Printf("the value at index %d and the value is %d\n", index, value)
	// }

	// // how we can take input in the slice and how we will print that 
	// var n int
	// fmt.Println("enter the value of n ")
	// fmt.Scan(&n)

	// name := []string{}

	// for i := 0; i < n; i++ {
	// 	var temp string
	// 	fmt.Printf("enter name \n")
	// 	fmt.Scan(&temp)
	// 	name = append(name, temp)
	// }

	// for index, value := range name {
	// 	fmt.Printf("entered name at index %d is %s\n", index, value)

	// }

	// running like a while loop by appying some condition on that
	x:=0
	for{
		fmt.Println("infinite loop")
		if(x>5){
			break
		}
		x++;
	}






// 	The for loop in Go is versatile and can be used to perform a wide 
// range of operations. Here are the different types of operations and 
// variations you can do with it:

// 1. Basic Iteration (C-style):
// Iterate over a range of numbers using an index.


for i := 0; i < 10; i++ {
    fmt.Println(i)
}



// 2. While-Like Loop:
// Simulate a while loop by using a single condition.


i := 0
for i < 10 {
    fmt.Println(i)
    i++
}


// 3. Infinite Loop:
// Run a loop indefinitely until explicitly broken.


for {
    fmt.Println("This will run forever unless stopped!")
    break // Prevent infinite loop
}



// 4. Range Loop:
// Iterate over slices, arrays, strings, maps, or channels.

// Slice or Array:

nums := []int{1, 2, 3, 4, 5}
for index, value := range nums {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}



// String (Character Iteration):

str := "hello"
for index, char := range str {
    fmt.Printf("Index: %d, Character: %c\n", index, char)
}


// Map:

m := map[string]int{"a": 1, "b": 2}
for key, value := range m {
    fmt.Printf("Key: %s, Value: %d\n", key, value)
}


// Channel:

ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
close(ch)

for value := range ch {
    fmt.Println(value)
}


// 5. Nested Loops:
// Use loops inside loops for multi-dimensional structures.


for i := 0; i < 3; i++ {
    for j := 0; j < 2; j++ {
        fmt.Printf("i: %d, j: %d\n", i, j)
    }
}

// 6. Break and Continue:
// Control the flow within loops.

// Break:

for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
    fmt.Println(i)
}


// Continue:

for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue
    }
    fmt.Println(i)
}


// 7. Using Labels:
// Break or continue nested loops using labels.


outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == j {
            continue outer
        }
        fmt.Printf("i: %d, j: %d\n", i, j)
    }
}


// 8. Iterating Backwards:
// Iterate in reverse order.


for i := 10; i >= 0; i-- {
    fmt.Println(i)
}

// 9. Dynamic Conditions:
// Use dynamic conditions to control loop execution.


// i := 0
// for i*i < 100 {
//     fmt.Println(i)
//     i++
// }

// 10. Iterating Over Struct Fields (Indirectly):
// Use reflection to iterate over fields of a struct.


// import 

type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 25}
val := reflect.ValueOf(p)
for i := 0; i < val.NumField(); i++ {
    fmt.Println(val.Field(i))
}


}








