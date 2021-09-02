package main

// The four adjacent digits in the 1000-digit number that have the greatest product are 9 x 9 x 8 x 9 = 5832
// Find the n adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?

func LargestProduct(ary []int, size int) int {
	left := 0
	right := size - 1
	max := 0

	for {
		if right >= len(ary) {
			return max
		}

		product := 1
		for i := left; i <= right; i++ {
			product *= ary[i]
		}

		if max < product {
			max = product
		}

		left++
		right++
	}
}
