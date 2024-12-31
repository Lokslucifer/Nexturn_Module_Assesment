package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	id         int
	name       string
	age        int
	department string
}

var departments = []string{"HR", "IT"}
var employees []Employee

func AddEmployee(emp Employee) error {
	if emp.age < 18 {
		return errors.New("employee is too young")
	}
	for _, e := range employees {
		if emp.id == e.id {
			return errors.New("id already exists")
		}
	}
	employees = append(employees, emp)
	return nil
}

func SearchEmployee(id int) (Employee, error) {
	for _, e := range employees {
		if id == e.id {
			return e, nil
		}
	}
	return Employee{}, errors.New("id does not exist")
}

func ListEmployees() {
	if len(employees) == 0 {
		fmt.Println("No employees to display.")
		return
	}
	for _, e := range employees {
		fmt.Printf("ID: %d, Name: %s, Age: %d, Department: %s\n", e.id, e.name, e.age, e.department)
	}
}

func FilterEmployees(depart string) []Employee {
	var filtered []Employee
	for _, e := range employees {
		if e.department == depart {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

func CountEmployees() map[string]int {
	counts := make(map[string]int)
	for _, dep := range departments {
		count := 0
		for _, e := range employees {
			if e.department == dep {
				count++
			}
		}
		counts[dep] = count
	}
	return counts
}

func main() {
	var choice int
	fmt.Println("The commands:")
	fmt.Println("1: Add Employee")
	fmt.Println("2: Search Employee")
	fmt.Println("3: List Employees by Department")
	fmt.Println("4: Count Employees")
	fmt.Println("5: Exit")

	for {
		fmt.Println("\nEnter your choice:")
		fmt.Scanf("%d", &choice)
		fmt.Scanln()
		switch choice {
		case 1:
			var emp Employee
			var deptid int
			fmt.Println("Enter Employee ID, Name, and Age:")
			fmt.Scanf("%d %s %d", &emp.id, &emp.name, &emp.age)
			fmt.Scanln()

			fmt.Println("Select department (0: HR, 1: IT):")
			fmt.Scanf("%d", &deptid)
			fmt.Scanln()
			if deptid >= 0 && deptid < len(departments) {
				emp.department = departments[deptid]
			} else {
				fmt.Println("Unknown selection")
				continue
			}
			err := AddEmployee(emp)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Employee added successfully.")
			}

		case 2:
			var id int
			fmt.Println("Enter Employee ID to search:")
			fmt.Scanf("%d", &id)
			fmt.Scanln()
			emp, err := SearchEmployee(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Found Employee: %+v\n", emp)
			}

		case 3:
			var deptid int
			fmt.Println("Select department to filter (0: HR, 1: IT):")
			fmt.Scanf("%d", &deptid)
			fmt.Scanln()
			if deptid >= 0 && deptid < len(departments) {
				department := departments[deptid]
				filtered := FilterEmployees(department)
				fmt.Printf("Employees in %s Department:\n", department)
				for _, e := range filtered {
					fmt.Printf("ID: %d, Name: %s, Age: %d\n", e.id, e.name, e.age)
				}
			} else {
				fmt.Println("Unknown selection")
			}

		case 4:
			counts := CountEmployees()
			fmt.Println("Employee Counts by Department:")
			for dep, count := range counts {
				fmt.Printf("%s: %d\n", dep, count)
			}

		case 5:
			fmt.Println("Exiting program.")
			return

		default:
			fmt.Println("Invalid command. Please try again.")
		}
	}
}
