package functions

func Min(numbers []int) int {
	// [4, 2, 20, 8, 11, 3, 10, 20, 11, 20]
	// output:
	// minimum number = 2

	lowest := 10

	for _, index := range numbers {
		if index < lowest {
			lowest = index
		}
	}

	return lowest
}

func Max(numbers []int) int {
	// [4, 2, 20, 8, 11, 3, 10, 20, 11, 20]
	// output:
	// maximum number = 20

	highest := 0

	for _, index := range numbers {
		if index > highest {
			highest = index
		}
	}

	return highest
}
