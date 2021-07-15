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

	var (
		women, men           [][]interface{}
		datasWomen, datasMen []string
	)

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
		datasMen = append(datasMen, employment(index[0], index[1]))
	}

	for _, index := range women {
		datasWomen = append(datasWomen, employment(index[0], index[1]))
	}

	fmt.Println("men employees = " + strings.Join(datasMen, ", "))
	fmt.Println("women employees = " + strings.Join(datasWomen, ", "))
}

func employment(name, year interface{}) string {
	workingYears := fmt.Sprintf("%v (%v years of employment)", name, year)

	return workingYears
}
