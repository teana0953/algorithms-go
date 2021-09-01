package main

func AveragePair(sortedAry []float32, target float32) (result [][]float32) {
	leftPointer := 0
	rightPointer := len(sortedAry) - 1

	for {
		if rightPointer <= leftPointer {
			break
		}

		currentValue := (sortedAry[leftPointer] + sortedAry[rightPointer]) / 2
		if currentValue > target {
			rightPointer = rightPointer - 1
		} else if currentValue < target {
			leftPointer = leftPointer + 1
		} else {
			result = append(result, []float32{sortedAry[leftPointer], sortedAry[rightPointer]})
			rightPointer = rightPointer - 1
			leftPointer = leftPointer + 1
		}
	}

	return result
}
