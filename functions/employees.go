package functions

import (
	"fmt"
	"log"
	"strings"
)

func Employees() {
	// [["employee ABC", 0, 3], ["employee DEF", 1, 5], ["employee GHI", 1, 1], ["employee JKL", 0, 10]]
	// output:
	// men employees = employee ABC (3 years of employment), employee JKL (10 years of employment)
	// women employees = employee DEF (5 years of employment), employee GHI (1 years of employment)

	var women, men []string

	employee := [][]interface{}{
		{"employee ABC", 0, 3}, {"employee DEF", 1, 5}, {"employee GHI", 1, 1}, {"employee JKL", 0, 10},
	}

	for _, index := range employee {
		switch index[1] {
		case 0:
			men = append(men, employment(index[0], index[2]))
		case 1:
			women = append(women, employment(index[0], index[2]))
		default:
			log.Fatalln("No employee gender")
		}
	}

	fmt.Println("men employees = " + strings.Join(men, ", "))
	fmt.Println("women employees = " + strings.Join(women, ", "))
}

func employment(name, year interface{}) string {
	workingYears := fmt.Sprintf("%v (%v years of employment)", name, year)

	return workingYears
}
