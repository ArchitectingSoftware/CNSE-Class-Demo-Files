package main

import (
	"fmt"

	"architectingsoftware.com/gotutorial/demopackage"
)

//This demo shows some unique aspects of packages in go.  It uses
//the student.go file in the demopackage directory.

func basicPackageDemo() {
	//This is a basic package demo.  We are importing a package from the
	//architectingsoftware.com/gotutorial/demopackage directory.  This is a
	//package that I created for this tutorial.  It is not a standard go package.
	//You can see the code for this package in the demopackage directory.
	s := demopackage.New()

	//Notice that we can easily access the exported fields (the ones that
	//start with a capital letter) of the struct.  We cannot access the
	//unexported fields (the ones that start with a lowercase letter).
	fmt.Println("NAME:", s.Name)
	fmt.Println("YEAR:", s.Year)
	fmt.Println("GPA:", s.GPA)

	//We can also call the exported methods of the struct, but not the
	//unexported methods. Thus we cannot see the SSN of the student,
	//or the tuition balance.  In order to operate on the unexported
	//fields we need to use exported receiver methods.
	fmt.Println("SSN:", s.GetSSN("secret"))
	fmt.Println("Adjust Balance: ", "New Balance is:", s.AdjustBalance(100.0))

}

func RunPackagesDemo() {
	fmt.Println("------ Running Packages Demo ------")
	basicPackageDemo()
	fmt.Printf("-----------------------------------\n\n")
}
