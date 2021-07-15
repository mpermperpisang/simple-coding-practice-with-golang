package functions

import "fmt"

func FibonacciNumbers(number int) []string {
	// There are N numbers of Fibonacci
	// output:
	// fibonacci numbers = [0, 1, 1, 2, 3, 5, 8, 13, 21, 34]

	var (
		numbers   []int
		fibonacci []string
	)

	index := 2
	numbers = append(numbers, 0, 1)

	for index < number {
		numbers = append(numbers, numbers[index-1]+numbers[index-2])
		index += 1
	}

	for _, index := range numbers {
		fibonacci = append(fibonacci, fmt.Sprintf("%v", index))
	}

	return fibonacci
}
