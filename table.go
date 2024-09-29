package main

import "fmt"

func Table(n int) {
	fmt.Println("Multiplication table:")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}
