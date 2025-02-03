package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	fmt.Println("learing channels...")
// 	mych:=make(chan int)
// 	mych <- 5;
// 	fmt.Println(<-mych)
// }

// ....Agar koi bhi receiver available nahi hai, toh program deadlock me chala jata hai.
// ....Agar sender ko block hone se bachana hai, toh ya to goroutine use karo ya buffered channel lo.
// Channel ka Basic Rule
// Agar koi ek goroutine channel par data send kar raha hai,
// toh wo tab tak block rahega jab tak koi usko receive na kare.
// Agar koi ek goroutine channel se data receive kar raha hai,
//  toh wo tab tak block rahega jab tak koi usme data send na kare.

// func main() {
// 	fmt.Println("learing channels...")
// 	mych:=make(chan int)
// 	wg:=&sync.WaitGroup{}
// 	// mych <- 5;
// 	// fmt.Println(<-mych)

// 	wg.Add(2)
// 	// receive channel
// 	go func(ch chan int, wg *sync.WaitGroup){
// 		fmt.Println(<-ch)
// 		wg.Done()
// 	}(mych, wg)

// 	// send channel
// 	go func(ch chan int, wg *sync.WaitGroup){
// 		mych <- 5
// 		wg.Done()
// 	}(mych, wg)
// 	wg.Wait()
// }

// func main() {
// 	fmt.Println("learing channels...")
// 	mych:=make(chan int)
// 	wg:=&sync.WaitGroup{}
// 	// mych <- 5;
// 	// fmt.Println(<-mych)

// 	wg.Add(2)
// 	// receive channel
// 	go func(ch chan int, wg *sync.WaitGroup){
// 		fmt.Println(<-ch)
// 		wg.Done()
// 	}(mych, wg)

// 	// send channel
// 	go func(ch chan int, wg *sync.WaitGroup){
// 		mych <- 5
// 		mych <- 6
// 		wg.Done()
// 	}(mych, wg)
// 	wg.Wait()
// }

// In the above case i am filling the channel with two data.
// while only one time i am receiveing...that means it will also show deadlock.
// for resolveing this issue i will use buffer size like mych:=make(chan int,1)

// func main() {
// 	fmt.Println("learing channels...")
// 	mych:=make(chan int,1)
// 	wg:=&sync.WaitGroup{}
// 	// mych <- 5;
// 	// fmt.Println(<-mych)

// 	wg.Add(2)
// 	// receive channel
// 	go func(ch chan int, wg *sync.WaitGroup){
// 		fmt.Println(<-ch)
// 		wg.Done()
// 	}(mych, wg)

// 	// send channel
// 	go func(ch chan int, wg *sync.WaitGroup){
// 		mych <- 5
// 		mych <- 6
// 		wg.Done()
// 	}(mych, wg)
// 	wg.Wait()
// }

// func main() {
// 	fmt.Println("learing channels...")
// 	mych:=make(chan int)
// 	wg:=&sync.WaitGroup{}
// 	// mych <- 5;
// 	// fmt.Println(<-mych)

// 	wg.Add(2)
// 	// receive channel
// 	go func(ch chan int, wg *sync.WaitGroup){
// 		fmt.Println(<-ch)
// 		wg.Done()
// 	}(mych, wg)

// 	// send channel
// 	go func(ch chan int, wg *sync.WaitGroup){
// 		mych <- 5
// 		mych <- 6
// 		close(mych) // close channel................................................................
// 		wg.Done()
// 	}(mych, wg)
// 	wg.Wait()
// }

// Directional channels ka use sirf tab zaroori hota hai jab tum function ko sirf
// receiving ya sending ke liye restrict karna chahte ho.
// Agar function me general channel use ho raha hai, toh direct chan int bhi theek hai

// if i put close(mych)  in the recevier channel then that will run infinite time loop
// and it will not block the program until receiver channel is closed.
// for resolving this issue we will use <-chan and chan<- key

func main() {
	fmt.Println("Learning channels...")

	mych := make(chan int) // Unbuffered channel
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// Receive channel: Continuous listening until channel is closed
	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for val := range ch { // Loop until channel is closed
			fmt.Println("Received:", val)
		}
		fmt.Println("Receiver Goroutine Exiting...")
	}(mych, wg)

	// Send channel: Sends multiple values and then closes the channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		ch <- 5
		ch <- 6
		ch <- 7
		close(ch) // Closing channel after sending all values
		fmt.Println("Sender Goroutine Exiting...")
	}(mych, wg)

	wg.Wait()
	fmt.Println("Main function Exiting...")
}
