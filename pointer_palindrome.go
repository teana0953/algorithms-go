package main

import "strings"

func IsPalindrome(str string) bool {
	strAry := strings.Split(str, "")

	leftPointer := 0
	rightPointer := len(strAry) - 1

	for {
		if leftPointer >= rightPointer {
			return true
		}

		if strAry[leftPointer] != strAry[rightPointer] {
			return false
		}

		leftPointer = leftPointer + 1
		rightPointer = rightPointer - 1
	}
}
