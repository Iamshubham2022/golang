package main

import "fmt"

// Corrected simplefunction
func simplefunction(a, b int) int {
    fmt.Println("simple function")
    return a + b
}

// func simplefunction(a int , b ) int { // this shows error because of (a int ,b)
//     fmt.Println("simple function")
//     return a + b
// }

// Corrected simplefun
func simplefun(a, b int) (result int) {
    fmt.Println("simple function")
    result = a + b
    return // Named return variable "result" is returned automatically
}

func main() {
    fmt.Println("function start now")
    ans := simplefunction(10, 5)
    fmt.Println("The addition of two numbers is:", ans)

    an := simplefun(10, 5)
    fmt.Println("The addition of two numbers is:", an)
}
