package main

func reverseString(s string) string {
	revstr := []rune(s)
	for i, j := 0, len(revstr)-1; i < j; i, j = i+1, j-1 {
		revstr[i], revstr[j] = revstr[j], revstr[i]
	}
	return string(revstr)
}
