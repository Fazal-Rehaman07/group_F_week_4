package main

import "fmt"

func fibonacci(n int) {
	a := 0
	b := 1
	nxt := a + b
	fmt.Printf("%d %d ", a, b) //print the first two numbers in the Fibonacci series

	for i := 3; i <= n; i++ {
		fmt.Print(nxt, " ") //print the next number in the series
		a = b 
		b = nxt
		nxt = a + b //calculate the next number in the series
	}

	fmt.Println() 
}
