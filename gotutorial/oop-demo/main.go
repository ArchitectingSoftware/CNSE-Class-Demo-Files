package main

import (
	"fmt"

	oopobjects "architectingsoftware.com/gotutorial/oop-demo/oop-objects"
)

func main() {

	prof := oopobjects.GetProfessor()
	student := oopobjects.GetStudent()
	employee := oopobjects.GetEmployee()

	fmt.Println("Professor Name: ", prof.Name)
	fmt.Println("Student Name: ", student.Name)
	fmt.Println("Employee Name: ", employee.Name)

	fmt.Println("Professor Age: ", prof.GetAge())
	fmt.Println("Student Age: ", student.GetAge())
	fmt.Println("Employee Age: ", employee.GetAge())

	fmt.Println("Student GPA: ", student.GPA)

	fmt.Println("Professor Department: ", prof.Department)

	fmt.Println("Employee Title: ", employee.Title)
	fmt.Println("Professor Title: ", prof.Title)

	oldSalary, newSalary := prof.GiveRaise(0.1)
	fmt.Println("Give Prof a Raise, Before: ", oldSalary, " after: ", newSalary)
}
