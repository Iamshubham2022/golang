package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Is_Adult bool    `json:"is_adult"`
}

type Student struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	RollNo int    `json:"roll_no"`
}

func main() {
	person :=Person{Name:"John",Age :21,Is_Adult :true}
	student :=Student{Name:"shubham", Age :12, RollNo :2101066 }

	// Marshaling the struct to JSON format...
	// Note: Marshal function converts the given struct into a JSON object.
    // It returns the JSON formatted byte slice.
   
	jsonData,err:=json.Marshal(person)
	if err!=nil{
		fmt.Println("Error marshaling person",err)
        return
	}
	fmt.Println("JSON data is saved as json", string(jsonData))

	studData,error:=json.Marshal(student)
	if error!=nil{
		fmt.Println("Error marshaling person",error)
        return
	}
	fmt.Println("JSON data is saved as json", string(studData))

	// Unmarshaling JSON data to a struct...
	// Note: Unmarshal function converts the JSON object into a given struct.
    // It returns the number of fields successfully unmarshaled.
   

	var newPerson Person
	var newStudent Student

	errors:=json.Unmarshal(jsonData,&newPerson)
	if errors!=nil{
        fmt.Println("Error unmarshalling person",errors)
        return
    }
	fmt.Println("Unmarshaled person is",newPerson)

	er:=json.Unmarshal(studData,&newStudent)
	if er!=nil{
        fmt.Println("Error unmarshalling student",er)
        return
    }
	fmt.Println("Unmarshaled student is",newStudent)


}
