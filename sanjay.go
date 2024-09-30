package main

import (
	"fmt"
)

func decimalToBinary(n int) string {
	if n == 0 {
		return "0"
	}

	var binary string
	for n > 0 {
		remainder := n % 2
		binary = fmt.Sprintf("%d%s", remainder, binary)
		n /= 2
	}
	return binary
}
