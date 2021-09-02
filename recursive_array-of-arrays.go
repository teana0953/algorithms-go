package main

// Write a function that collects all value in an array of arrays and return an array of values collected
// ex. [[[["a", [["b", ["c"]],["d"]]], [["e"]],[[["f", "g", "h"]]]]]] => [a, b, c, d, e, f, g, h]

func ArrayOfArrays(ary []interface{}, acc []string) []string {
	if acc == nil {
		acc = []string{}
	}

	for _, item := range ary {
		if value, ok := item.([]interface{}); ok {
			acc = ArrayOfArrays(value, acc)
		} else {
			value, _ := item.(string)
			acc = append(acc, value)
		}
	}

	return acc
}
