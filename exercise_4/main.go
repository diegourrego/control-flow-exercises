package exercise_4

import "fmt"

var employees = map[string]uint8{
	"Benjamin": 20,
	"Nahuel":   26,
	"Brenda":   19,
	"Dario":    44,
	"Pedro":    30,
}

func FindEmployee(nameEmployee string) (name string, age uint8, err error) {

	ageFound, ok := employees[nameEmployee]
	if !ok {
		fmt.Printf("Employee %s not found \n", nameEmployee)
		return "", 0, fmt.Errorf("employees[nameEmployee]: %w", err)
	}

	return nameEmployee, ageFound, nil
}

func EmployeesGreaterThan21() {
	var adultEmployees uint8

	for _, employeeAge := range employees {
		if employeeAge < 21 {
			continue
		}
		adultEmployees++
	}

	fmt.Printf("The company has %v employees up to 21\n", adultEmployees)

}

func AddNewEmployee(employeeName string, employeeAge uint8) {
	employees[employeeName] = employeeAge
	fmt.Println("Employees list:", employees)
}

func DeleteEmployee(employeeName string) {
	delete(employees, employeeName)
	fmt.Println("Employees list:", employees)
}
