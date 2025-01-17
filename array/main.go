package main

import "fmt"

func main(){
	  
	var number [5]int 

	var amount = [5]int{1,2,3,4,5} 

	// amount := [5]int{1,2,3,4,5}  if we not begin any variable name 
	// 							    with var then we use : (collon)

	fmt.Println(amount)
	fmt.Println(amount[4])
	fmt.Println(number)

	fmt.Println("length of array is : ", len(amount))
	fmt.Printf("length of array is : %d\n", len(amount))

	
}