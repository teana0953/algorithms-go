package main

import "strings"

// order matters
func IsSubsequence(newStr string, oldStr string) bool {
	if len(newStr) == 0 {
		return true
	}

	if len(newStr) > len(oldStr) {
		return false
	}

	newStrLeftPointer := 0
	newStrEndIndex := len(newStr) - 1
	newStrAry := strings.Split(newStr, "")

	oldStrLeftPointer := 0
	oldStrEndIndex := len(oldStr) - 1
	oldStrAry := strings.Split(oldStr, "")

	for {
		// end condition
		if newStrLeftPointer >= newStrEndIndex {
			return true
		} else if oldStrLeftPointer >= oldStrEndIndex {
			return false
		}

		newCurrent := newStrAry[newStrLeftPointer]
		oldCurrent := oldStrAry[oldStrLeftPointer]

		if newCurrent == oldCurrent {
			newStrLeftPointer = newStrLeftPointer + 1
			oldStrLeftPointer = oldStrLeftPointer + 1
		} else {
			oldStrLeftPointer = oldStrLeftPointer + 1
		}

	}
}
