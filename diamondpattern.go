package main

import "fmt"

func Diamond() {

	// Upper part of the diamond
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 6-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Lower part of the diamond
	for i := 6 - 1; i >= 1; i-- {
		for j := 1; j <= 6-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
