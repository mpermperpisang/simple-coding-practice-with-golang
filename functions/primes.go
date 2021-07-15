package functions

import "fmt"

func PrimesNumbers(numbers []int) []string {
	// [4, 2, 20, 8, 11, 3, 10, 20, 11, 20]
	// output:
	// primes numbers = [2, 3, 5, 7, 11, 13, 17, 19]

	var primes []string

	SortLowestNumbers(numbers)

	for index := numbers[0]; index <= numbers[len(numbers)-1]; index++ {
		prime := 0

		for num := 1; num <= index; num++ {
			if index%num == 0 {
				prime += 1
			}
		}

		if prime == 2 {
			primes = append(primes, fmt.Sprintf("%v", index))
		}
	}

	return primes
}
