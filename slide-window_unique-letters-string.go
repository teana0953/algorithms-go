package main

import "strings"

// Write a function called uniqueLetterString, which accepts a string and returns the length of the longest substring with all distinct characters
// ex. thisisshowwedoit => 6

func UniqueLetterString(str string) int {
	if len(str) == 0 {
		return 0
	}

	strAry := strings.Split(str, "")

	max := 0
	startPointer := 0
	endPointer := 0
	counter := map[string]bool{}

	for {
		if endPointer >= len(strAry) {
			return max
		}

		endCharacter := strAry[endPointer]
		if counter[endCharacter] {
			counter[strAry[startPointer]] = false
			startPointer++
		} else {
			counter[endCharacter] = true
			endPointer++

			// check whether currentLength is the max
			if (endPointer - startPointer) > max {
				max = endPointer - startPointer
			}
		}

	}
}
