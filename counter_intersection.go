package main

func FindIntersection(ary1 []int, ary2 []int) (results []int) {
	ary3 := append(ary1, ary2...)
	counter := map[int]int{}

	for _, item := range ary3 {
		counter[item] = counter[item] + 1
	}

	for key, value := range counter {
		if value >= 2 {
			results = append(results, key)
		}
	}

	return results
}
