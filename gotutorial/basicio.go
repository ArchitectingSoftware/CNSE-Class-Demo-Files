package main

import (
	"fmt"
	"log"
)

// notice the function begins with a lower case letter, you cannot call it outside of this
// package
func writeToStdOutput() {

	//Here is basic Println
	fmt.Println("Hello, World!")

	//String things together
	today := "Monday"          //Type assertion
	var favoriteNumber int = 6 //Specific definition of type with initialization
	fmt.Println("Today is", today, "and my favorite number is", favoriteNumber)

	//manually creating a type
	var classCode string
	classCode = "CS999"
	totalStudents := 30
	fmt.Printf("I can also do C style printf: Class = %s NumStudents = %d\n",
		classCode, totalStudents)
}

func builtInLogging() {
	//Go has a basic logging package
	courseName := "Cloud Native Software Engineering"
	log.Println("log can do Println, Printf, and Print")
	log.Printf("Hi, I hope you enjoy %s\n", "Go")
	log.Print("And this ", "class as well - ", courseName)
}

func RunBasicIODemo() {
	fmt.Println("------ Running Basic IO Demo ------")
	writeToStdOutput()
	builtInLogging()
	fmt.Printf("-----------------------------------\n\n")
}
