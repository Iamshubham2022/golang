package main

import "fmt"

var person struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Is_Adult int `json:"is_adult"`
}

var student struct{
	Name string `json:"name"`
	Age int `json:"age"`
	RollNo int `json:"rollno"`
}

func main(){
	fmt.Println("we are learing golang")

}