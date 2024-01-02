package main

import (
	"controlFlow_exercises/exercise_1"
	"controlFlow_exercises/exercise_2"
	"controlFlow_exercises/exercise_3"
	"controlFlow_exercises/exercise_4"
	"fmt"
)

func main() {
	exercise_1.FirstExercise("test")
	exercise_2.ExerciseTwo(32, "employee", 2, 250320.67)
	fmt.Printf("Mes: %s \n", exercise_3.ExerciseThree(5))

	nameEmployee, ageEmployee, err := exercise_4.FindEmployee("Benjamin")

	if err != nil {
		formattedError := fmt.Errorf("FindEmployee(): %w", err)
		fmt.Println(formattedError)
		return
	}

	fmt.Printf("The employee %s has %v years old\n", nameEmployee, ageEmployee)

	exercise_4.EmployeesGreaterThan21() // 3
	exercise_4.AddNewEmployee("Federico", 25)
	exercise_4.EmployeesGreaterThan21() // 4
	exercise_4.DeleteEmployee("Pedro")
}
