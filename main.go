package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simple-coding-practice-with-golang/functions"
)

func init() {
	fmt.Println("This is just a simple quiz.")
	fmt.Println("===========================")
}

func main() {
	numbers := []int{4, 2, 20, 8, 11, 3, 10, 20, 11, 20}

	fmt.Println("minimum number = " + strconv.Itoa(functions.Min(numbers)))
	fmt.Println("maximum number = " + strconv.Itoa(functions.Max(numbers)))
	fmt.Println("---------------------------")
	fmt.Println("odd numbers = [" + joinArray(functions.OddNumbers(numbers)))
	fmt.Println("even numbers = [" + joinArray(functions.EvenNumbers(numbers)))
	fmt.Println("primes numbers = [" + joinArray(functions.PrimesNumbers(numbers)))
	fmt.Println("fibonacci numbers = [" + joinArray(functions.FibonacciNumbers(10)))
	fmt.Println("---------------------------")
	fmt.Println("sort lowest - highest = [" + joinArray(functions.SortLowestNumbers(numbers)))
	fmt.Println("sort highest - lowest = [" + joinArray(functions.SortHighestNumbers(numbers)))
	fmt.Println("---------------------------")
	fmt.Println("mean = " + strconv.Itoa(functions.Mean(numbers)))
	fmt.Println("median = " + strconv.Itoa(functions.Median(numbers)))
	fmt.Println("modus = " + strconv.Itoa(functions.Modus(numbers)))
	fmt.Println("---------------------------")
	functions.Employees()
}

func joinArray(array []string) string {
	return strings.Join(array, ", ") + "]"
}
