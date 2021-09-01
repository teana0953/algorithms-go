package main

import "strings"

func SameFrequency(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	counter1 := map[string]int{}
	counter2 := map[string]int{}

	str1Ary := strings.Split(str1, "")
	str2Ary := strings.Split(str2, "")

	for _, item := range str1Ary {
		counter1[item] = counter1[item] + 1
	}

	for _, item := range str2Ary {
		counter2[item] = counter2[item] + 1
	}

	for key, value := range counter1 {
		if counter2[key] != value {
			return false
		}
	}

	return true
}
