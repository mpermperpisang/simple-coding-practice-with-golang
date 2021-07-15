package functions

import "fmt"

func SortLowestNumbers(numbers []int) []string {
	// [4, 2, 20, 8, 11, 3, 10, 20, 11, 20]
	// output:
	// sort lowest - highest = [2, 3, 4, 8, 10, 11, 11, 20, 20, 20]

	var sortLowestNumber []string

	for index := 0; index < len(numbers); index++ {
		for index2 := 0; index2 < len(numbers); index2++ {
			if numbers[index] < numbers[index2] {
				temp := numbers[index]
				numbers[index] = numbers[index2]
				numbers[index2] = temp
			}
		}
	}

	for index := 0; index < len(numbers); index++ {
		sortLowestNumber = append(sortLowestNumber, fmt.Sprintf("%v", numbers[index]))
	}

	return sortLowestNumber
}

func SortHighestNumbers(numbers []int) []string {
	// [4, 2, 20, 8, 11, 3, 10, 20, 11, 20]
	// output:
	// sort highest - lowest = [20, 20, 20, 11, 11, 10, 8, 4, 3, 2]

	var sortHighestNumber []string

	for index := 0; index < len(numbers); index++ {
		for index2 := 0; index2 < len(numbers); index2++ {
			if numbers[index] > numbers[index2] {
				temp := numbers[index]
				numbers[index] = numbers[index2]
				numbers[index2] = temp
			}
		}
	}

	for index := 0; index < len(numbers); index++ {
		sortHighestNumber = append(sortHighestNumber, fmt.Sprintf("%v", numbers[index]))
	}

	return sortHighestNumber
}
