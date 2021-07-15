package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func init() {
	fmt.Println("This is just a simple quiz.")
	fmt.Println("===========================")
}

func main() {
	numbers := []int{4, 2, 20, 8, 11, 3, 10, 20, 11, 20}

	fmt.Println("minimum number = " + strconv.Itoa(min(numbers)))
	fmt.Println("maximum number = " + strconv.Itoa(max(numbers)))
	fmt.Println("---------------------------")
	fmt.Println("odd numbers = [" + strings.Join(oddNumbers(numbers), ", ") + "]")
	fmt.Println("even numbers = [" + strings.Join(evenNumbers(numbers), ", ") + "]")
	fmt.Println("primes numbers = [" + strings.Join(primesNumbers(numbers), ", ") + "]")
	fmt.Println("fibonacci numbers = [" + strings.Join(fibonacciNumbers(10), ", ") + "]")
	fmt.Println("---------------------------")
	fmt.Println("sort lowest - highest = [" + strings.Join(sortLowestNumbers(numbers), ", ") + "]")
	fmt.Println("sort highest - lowest = [" + strings.Join(sortHighestNumbers(numbers), ", ") + "]")
	fmt.Println("---------------------------")
	fmt.Println("mean = " + strconv.Itoa(mean(numbers)))
	fmt.Println("median = " + strconv.Itoa(median(numbers)))
	fmt.Println("modus = " + strconv.Itoa(modus(numbers)))
	fmt.Println("---------------------------")
	employees()
}

func min(numbers []int) int {
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

func max(numbers []int) int {
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

func oddNumbers(numbers []int) []string {
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

func evenNumbers(numbers []int) []string {
	// [4, 2, 20, 8, 11, 3, 10, 20, 11, 20]
	// output:
	// even numbers = [4, 2, 20, 8, 10, 20, 20]

	var even []string

	for index := 0; index < len(numbers); index++ {
		if numbers[index]%2 == 0 {
			even = append(even, fmt.Sprintf("%v", numbers[index]))
		}
	}

	return even
}

func primesNumbers(numbers []int) []string {
	// [4, 2, 20, 8, 11, 3, 10, 20, 11, 20]
	// output:
	// primes numbers = [2, 3, 5, 7, 11, 13, 17, 19]

	var primes []string

	sortLowestNumbers(numbers)

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

func fibonacciNumbers(number int) []string {
	// There are N numbers of Fibonacci
	// output:
	// fibonacci numbers = [0, 1, 1, 2, 3, 5, 8, 13, 21, 34]

	var (
		numbers   []int
		fibonacci []string
	)
	numbers = append(numbers, 0, 1)
	index := 2

	for index < number {
		numbers = append(numbers, numbers[index-1]+numbers[index-2])
		index += 1
	}

	for _, index := range numbers {
		fibonacci = append(fibonacci, fmt.Sprintf("%v", index))
	}

	return fibonacci
}

func sortLowestNumbers(numbers []int) []string {
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

func sortHighestNumbers(numbers []int) []string {
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

func mean(numbers []int) int {
	// [4, 2, 20, 8, 11, 3, 10, 20, 11, 20]
	// output:
	// mean = 10

	var number int

	for _, index := range numbers {
		number += index
	}

	return number / (len(numbers))
}

func median(numbers []int) int {
	// [4, 2, 20, 8, 11, 3, 10, 20, 11, 20]
	// output:
	// median = 11

	sortLowestNumbers(numbers)

	return numbers[len(numbers)/2]
}

func modus(numbers []int) int {
	// [4, 2, 20, 8, 11, 3, 10, 20, 11, 20]
	// output:
	// modus = 20

	var modusNumber [][]int
	number, highest, modus := 0, 0, 0

	for index := 0; index < len(numbers)-1; index++ {
		if numbers[index+1] == numbers[index] {
			number += 1
		} else {
			number = 1
		}

		if index == 0 {
			modusNumber = append(modusNumber, []int{numbers[0], number})
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

func employees() {
	// [["employee ABC", 0, 3], ["employee DEF", 1, 5], ["employee GHI", 1, 1], ["employee JKL", 0, 10]]
	// output:
	// men employee = employee ABC (3 years of employment), employee JKL (10 years of employment)
	// women employee = employee DEF (5 years of employment), employee GHI (1 years of employment)

	var women, men [][]interface{}
	var datasWomen, datasMen []string
	employee := [][]interface{}{
		{"employee ABC", 0, 3}, {"employee DEF", 1, 5}, {"employee GHI", 1, 1}, {"employee JKL", 0, 10},
	}

	for _, index := range employee {
		switch index[1] {
		case 0:
			men = append(men, []interface{}{index[0], index[2]})
		case 1:
			women = append(women, []interface{}{index[0], index[2]})
		default:
			log.Fatalln("No employee gender")
		}
	}

	for _, index := range men {
		workingYears := " (" + fmt.Sprintf("%v", index[1]) + " years of employment)"
		datasMen = append(datasMen, fmt.Sprintf("%v", index[0])+workingYears)
	}

	for _, index := range women {
		workingYears := " (" + fmt.Sprintf("%v", index[1]) + " years of employment)"
		datasWomen = append(datasWomen, fmt.Sprintf("%v", index[0])+workingYears)
	}

	fmt.Println("men employee = " + strings.Join(datasMen, ", "))
	fmt.Println("women employee = " + strings.Join(datasWomen, ", "))
}
