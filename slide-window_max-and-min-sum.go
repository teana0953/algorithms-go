package main

import "math"

// Write two functions that calculate the max and min sum of n consecutive numbers in an array

// to arr.length - 1 - (n - 1) => arr.length - n

func MaxSum(ary []int, size int) *int {
	if size > len(ary) {
		return nil
	}

	max := int(math.Inf(-1))

	for i := 0; i <= len(ary)-size; i++ {
		attempt := 0
		for j := i; j < i+size; j++ {
			attempt = attempt + ary[j]
		}

		if attempt > max {
			max = attempt
		}
	}

	return &max
}

func MaxSumImproved(ary []int, size int) *int {
	if size > len(ary) {
		return nil
	}

	max := 0

	// first window
	for i := 0; i < size; i++ {
		max = max + ary[i]
	}

	temp := max
	for i := size; i < len(ary); i++ {
		// add next and minus first of last window
		temp = temp + ary[i] - ary[i-size]

		if temp > max {
			max = temp
		}
	}

	return &max
}

func MinSum(ary []int, size int) *int {
	if size > len(ary) {
		return nil
	}

	min := 0

	for i := 0; i < size; i++ {
		min += ary[i]
	}

	temp := min
	for i := size; i < len(ary); i++ {
		temp = temp + ary[i] - ary[i-size]
		if temp < min {
			min = temp
		}
	}

	return &min
}
