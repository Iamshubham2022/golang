package main

import "fmt"

func main(){
	var month string 
	fmt.Println("enter the value of your month ")
	fmt.Scan(&month);
	
	//swich multiple value and basic
	switch month {
	case "jan", "feb", "mar":
		fmt.Printf("winter \n")
	case "jun" :
		fmt.Printf("summer \n")
	case "dec":
		fmt.Printf("snow \n")
	default:
		fmt.Printf("enter month \n")
	}
	
	
	var age int 
	fmt.Println("enter the value of your age ")
	fmt.Scan(&age);


	// switch expression 
	switch{
	case age==17:
		fmt.Print("you are not able to drive car")
	case  age <18:
		fmt.Print("you are able to drive car but next year")
	case age>=18:
		fmt.Print("you are able to drive car")
	default:
		fmt.Print("enter your age")
	}
	
}