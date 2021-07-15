package functions

func Mean(numbers []int) int {
	// [4, 2, 20, 8, 11, 3, 10, 20, 11, 20]
	// output:
	// mean = 10

	var number int

	for _, index := range numbers {
		number += index
	}

	return number / (len(numbers))
}

func Median(numbers []int) int {
	// [4, 2, 20, 8, 11, 3, 10, 20, 11, 20]
	// output:
	// median = 11

	SortLowestNumbers(numbers)

	return numbers[len(numbers)/2]
}

func Modus(numbers []int) int {
	// [4, 2, 20, 8, 11, 3, 10, 20, 11, 20]
	// output:
	// modus = 20

	var modusNumber [][]int
	number, highest, modus := 1, 0, 0
	modusNumber = append(modusNumber, []int{numbers[0], number})

	for index := 0; index < len(numbers)-1; index++ {
		if numbers[index] == numbers[index+1] {
			number += 1
		} else {
			number = 1
		}

		if numbers[index] != numbers[index+1] {
			modusNumber = append(modusNumber, []int{numbers[index+1], number})
		} else {
			modusNumber[len(modusNumber)-1][1] = number
		}
	}

	for index := 0; index < len(modusNumber); index++ {
		if modusNumber[index][1] > highest {
			modus = modusNumber[index][0]
			highest = modusNumber[index][1]
		}
	}

	return modus
}
