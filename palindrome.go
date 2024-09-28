package main

func palindrome(n int) bool {
	orig := n
	rev := 0

	for n > 0 {
		r := n % 10
		rev = rev * 10 + r
		n = n/10
	}

	return orig == rev
}