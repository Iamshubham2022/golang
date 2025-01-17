package main

import "fmt"

func main(){
	var age int
	fmt.Println("enter the variable value:",)
	fmt.Scan(&age)
	fmt.Printf("the given value is : %d\n",age);

	if(age>18){
		fmt.Print("you can drive car")
	} else if(age==17){
		fmt.Println("you can drive car but next year")
	} else{
		fmt.Printf("you are not able to drive a car because you are younger")
	}
}