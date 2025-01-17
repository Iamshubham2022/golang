package main

import (
	"fmt"
)

// slice is a dynamic array like vector in c++;
// to know data type of a variable use %T\n in Printf("data type is : %T\n",var)
func main(){
	// name := []string{"shubah","rohit","kohli"}
	// name = append(name, "jaiswal")
	// fmt.Printf("Data type of array is :%T\n",name)
	// fmt.Println("length of string :", len(name))
	// fmt.Println(name[3])
	// fmt.Println(name)

	// name := []int{}

	name:= make([]int,3,5)  //(array,length,capacity)
	name = append(name,2)
	name = append(name,3)

	// after slice capacity full  then slice dynamically increase their size; double 
	name = append(name,2)


	fmt.Println("slice:", name)
	fmt.Println("length:",len(name))
	fmt.Println("capacity:",cap(name))


}