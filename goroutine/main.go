package main

import (
	"fmt"
	"time"
)

func Addation(a, b int){
	fmt.Println(a + b)
	time.Sleep(1000*time.Millisecond)
	fmt.Println("hii this is from the additions")
}

func Subtraction(a, b int){
	fmt.Println(a - b)
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("hii this is from the subtractions")
	
}

func main() {
	fmt.Println("learing goroutine and checking how concurrency works and parallelism,so lets start")
	go Addation(5, 6)
	go Subtraction(7, 3)

	time.Sleep(2000 * time.Millisecond)

}
