package main

import "fmt"

func main() {
    fmt.Println("Welcome to Group F's Week 4 Project!\n")

    var choice int
    fmt.Println("Select the Function to Run")
    fmt.Println("1. Fibonacci Series")
    fmt.Println("2. Palindrome")
    fmt.Println("3. Diamond Pattern")

    fmt.Println("\nEnter Your Choice: ")
	fmt.Scan(&choice)
    
    switch choice{
    case 1:
        var n int
        fmt.Print("Enter the Limit (Positive No.) to Display Fibonacci Series: ")
	    fmt.Scan(&n)
        fmt.Printf("Fibonacci series till %d numbers:\n", n)
        fibonacci(n)

    case 2:
        var input int
        fmt.Print("Enter a number: ")
	    fmt.Scan(&input)
	    if palindrome(input) {
		fmt.Println("The number is a palindrome!")
	    } else {
		fmt.Println("The number is not a palindrome.")
	    }

    default:
        fmt.Println("Please try again by selecting a value between 1 and 5")

    }
    
    
    
}