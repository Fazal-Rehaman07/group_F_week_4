package main

import "fmt"

func fibonacci(n int) {
	a := 0
	b := 1
	nxt := a + b
	fmt.Printf("%d %d ", a, b)

	for i := 3; i <= n; i++ {
		fmt.Print(nxt, " ")
		a = b
		b = nxt
		nxt = a + b
	}

	fmt.Println()
}
