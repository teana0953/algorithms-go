package main

import (
	"math"
)

// Write a function called minSubLength which accepts two parameters, an array of positive integers and a positive integer. This function should return the minimal length of a continuous subarray - the sum of elements inside this subarray has to be greatter than or equal to the positive integer parameter. If subarray not found, then return 0.

// [8, 1, 6, 15, 3, 16, 5, 7, 14, 30, 12], 60 => [7, 14, 30, 12]

func MinSubLength(ary []int, sum int) int {
	startPointer := 0
	endPointer := 0
	total := 0

	isMinInit := false
	minLength := 0

	// lastIndex := len(ary) - 1

	for {
		if startPointer >= len(ary) {
			break
		}

		if total < sum && endPointer < len(ary) {
			total += ary[endPointer]
			endPointer = endPointer + 1
		} else if total >= sum {
			currentLength := endPointer - startPointer
			if !isMinInit {
				isMinInit = true
				minLength = currentLength
			} else if minLength > currentLength {
				minLength = currentLength
			}

			total = total - ary[startPointer]
			startPointer = startPointer + 1
		} else if endPointer >= len(ary) {
			break
		}
	}

	if minLength == int(math.Inf(1)) {
		return 0
	}

	return minLength
}
