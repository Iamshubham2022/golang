package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1. Basic Type Conversions
	var a int = 42
	var b float64 = float64(a) // Converting int to float64
	fmt.Printf("int to float64: %f\n", b)

	var c float64 = 3.14
	var d int = int(c) // Converting float64 to int (truncates decimal part)
	fmt.Printf("float64 to int: %d\n", d)

	// 2. String Conversions
	// Integer to string
	num := 123
	numStr := strconv.Itoa(num)
	fmt.Printf("Integer to string: %s\n", numStr)

	// String to integer
	str := "456"
	numFromStr, err := strconv.Atoi(str)
	if err == nil {
		fmt.Printf("String to integer: %d\n", numFromStr)
	} else {
		fmt.Println("Error converting string to integer:", err)
	}

	// Float to string
	floatNum := 3.14159
	floatStr := strconv.FormatFloat(floatNum, 'f', 2, 64) // 2 decimal places
	fmt.Printf("Float to string: %s\n", floatStr)

	// String to float
	strFloat := "123.456"
	floatFromStr, err := strconv.ParseFloat(strFloat, 64)
	if err == nil {
		fmt.Printf("String to float: %f\n", floatFromStr)
	} else {
		fmt.Println("Error converting string to float:", err)
	}

	// 3. Bool Conversions
	// Bool to string
	boolVal := true
	boolStr := strconv.FormatBool(boolVal)
	fmt.Printf("Bool to string: %s\n", boolStr)

	// String to bool
	strBool := "true"
	boolFromStr, err := strconv.ParseBool(strBool)
	if err == nil {
		fmt.Printf("String to bool: %t\n", boolFromStr)
	} else {
		fmt.Println("Error converting string to bool:", err)
	}

	// 4. Byte Slice to String and Vice Versa
	byteSlice := []byte("hello")
	strFromBytes := string(byteSlice) // Byte slice to string
	fmt.Printf("Byte slice to string: %s\n", strFromBytes)

	strToBytes := []byte("world") // String to byte slice
	fmt.Printf("String to byte slice: %v\n", strToBytes)

	// 5. Type Assertion (For interfaces)
	var i interface{} = "example string"
	strFromInterface, ok := i.(string) // Type assertion
	if ok {
		fmt.Printf("Type assertion to string: %s\n", strFromInterface)
	} else {
		fmt.Println("Type assertion failed")
	}

	



	// 8. Struct to JSON and Vice Versa (Using Encoding/JSON)
	type Person struct {
		Name string
		Age  int
	}
	// person := Person{Name: "Alice", Age: 30}

	// Struct to JSON
	// jsonBytes, err := json.Marshal(person)
	// if err == nil {
	// 	fmt.Printf("Struct to JSON: %s\n", jsonBytes)
	// }

	// // JSON to Struct
	// var personFromJSON Person
	// err = json.Unmarshal(jsonBytes, &personFromJSON)
	// if err == nil {
	// 	fmt.Printf("JSON to Struct: %+v\n", personFromJSON)
	// } else {
	// 	fmt.Println("Error converting JSON to struct:", err)
	// }
}
