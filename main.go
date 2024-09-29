package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("\nWelcome to Group F's Week 4 Project!")

	for {
		var choice int
		fmt.Println("\n\nSelect the Function to Run")
		fmt.Println("1. Fibonacci Series")
		fmt.Println("2. Palindrome")
		fmt.Println("3. Diamond Pattern")
		fmt.Println("4. Reverse String")
		fmt.Println("5. table")
		fmt.Println("6. Exit")

		fmt.Print("\nEnter Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var n int
			fmt.Print("Enter the Limit (Positive No.) to Display Fibonacci Series: ")
			fmt.Scan(&n)
			if n > 0 {
				fmt.Printf("\nFibonacci series till %d numbers:\n", n)
				fibonacci(n)
			} else {
				fmt.Printf("\nPlease try again with a Positive No. Greater than 0")
			}

		case 2:
			var input int
			fmt.Print("Enter a number: ")
			fmt.Scan(&input)
			if palindrome(input) {
				fmt.Println("\nThe number is a palindrome!")
			} else {
				fmt.Println("\nThe number is not a palindrome.")
			}

		case 3:
			Diamond()

		case 4:
			fmt.Print("Enter String to Reverse: ")
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\r')
			rev := reverseString(input)
			//reversed := reverseString("Hello World How are you")
			fmt.Println("Original - Reversed ", input)
			fmt.Println(rev)

		case 5:
			var input int
			fmt.Print("Enter a number: ")
			fmt.Scan(&input)
			Table(input)

		case 6:
			os.Exit(0)

		default:
			fmt.Println("\nPlease try again by selecting a value between 1 and 6")

		}
	}

}
