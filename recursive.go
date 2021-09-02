package main

func Recursive(n int) int {
	if n == 1 {
		return 10
	} else {
		return Recursive(n-1) + 15
	}
}
