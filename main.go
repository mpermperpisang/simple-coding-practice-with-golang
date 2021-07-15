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
	fmt.Println("odd numbers = [" + strings.Join(functions.OddNumbers(numbers), ", ") + "]")
	fmt.Println("even numbers = [" + strings.Join(functions.EvenNumbers(numbers), ", ") + "]")
	fmt.Println("primes numbers = [" + strings.Join(functions.PrimesNumbers(numbers), ", ") + "]")
	fmt.Println("fibonacci numbers = [" + strings.Join(functions.FibonacciNumbers(10), ", ") + "]")
	fmt.Println("---------------------------")
	fmt.Println("sort lowest - highest = [" + strings.Join(functions.SortLowestNumbers(numbers), ", ") + "]")
	fmt.Println("sort highest - lowest = [" + strings.Join(functions.SortHighestNumbers(numbers), ", ") + "]")
	fmt.Println("---------------------------")
	fmt.Println("mean = " + strconv.Itoa(functions.Mean(numbers)))
	fmt.Println("median = " + strconv.Itoa(functions.Median(numbers)))
	fmt.Println("modus = " + strconv.Itoa(functions.Modus(numbers)))
	fmt.Println("---------------------------")
	functions.Employees()
}
