package main

import "fmt"

func Check(stri string) bool {
	i := 0
	var s0, s0b, s1, s1b, s2, s2b int

	for i < len(stri) {
		if stri[i] == '(' {
			s0 = s0 + 1
		} else if stri[i] == ')' {
			s0b = s0b + 1
		} else if stri[i] == '[' {
			s1 = s1 + 1
		} else if stri[i] == ']' {
			s1b = s1b + 1
		} else if stri[i] == '{' {
			s2 = s2 + 1
		} else if stri[i] == '}' {
			s2b = s2b + 1
		}
		i++
	}

	if s0 == s0b && s1 == s1b && s2 == s2b {
		return true
	}
	return false
}

func main() {
	stri := "{(((({}))))}"
	fmt.Println(Check(stri))
}
