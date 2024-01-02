package exercise_2

import "fmt"

func ExerciseTwo(age uint8, employmentSituation string, yearsAntiquity int, salary float64) {
	switch {
	case age < 22, employmentSituation != "employee", yearsAntiquity < 1:
		fmt.Println("Denied loan")
	case salary < 100000:
		fmt.Println("Approved credit with interest")
	default:
		fmt.Println("Approved credit without interest")
	}
}
