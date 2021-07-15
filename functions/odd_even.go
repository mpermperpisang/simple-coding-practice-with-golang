package functions

import "fmt"

func OddNumbers(numbers []int) []string {
	// [4, 2, 20, 8, 11, 3, 10, 20, 11, 20]
	// output:
	// odd numbers = [11, 3, 11]

	var odd []string

	for _, index := range numbers {
		if index%2 == 1 {
			odd = append(odd, fmt.Sprintf("%v", index))
		}
	}

	return odd
}

func EvenNumbers(numbers []int) []string {
	// [4, 2, 20, 8, 11, 3, 10, 20, 11, 20]
	// output:
	// even numbers = [4, 2, 20, 8, 10, 20, 20]

	var even []string

	for _, index := range numbers {
		if index%2 == 0 {
			even = append(even, fmt.Sprintf("%v", index))
		}
	}

	return even
}
