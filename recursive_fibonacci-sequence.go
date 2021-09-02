package main

// Write a function that takes an integer N as an input and returns the Nth number in Fibonacci Sequence

func FibonacciSequence(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return FibonacciSequence(n-1) + FibonacciSequence(n-2)
	}
}
