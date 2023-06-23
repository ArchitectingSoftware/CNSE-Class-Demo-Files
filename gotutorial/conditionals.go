package main

import "fmt"

func basicConditionalIf() {
	//Go has basic conditionals
	var x int = 10
	if x > 5 {
		println("x is greater than 5")
	} else {
		println("x is less than or equal to 5")
	}
}

func basicConditionalSwitch() {
	//Go also has switch statements
	x := 10

	switch x {
	case 1:
		println("x is 1")
	case 2:
		println("x is 2")
	case 3:
		println("x is 3")
	default:
		println("x is not 1, 2, or 3")
	}

	//note, unlike C the go switch can work on different types
	//this is a string switch
	y := "hello"
	switch y {
	case "hello":
		println("y is hello")
	case "goodbye":
		println("y is goodbye")
	default:
		println("y is something else")
	}

	//GO can also switch on types
	//this is a type switch
	var z interface{}
	z = 1
	switch z.(type) {
	case int:
		println("z is an int")
	case string:
		println("z is a string")
	default:
		println("z is something else")
	}
}

func RunConditionalDemo() {
	fmt.Println("------ Running Conditional Demo ------")
	basicConditionalIf()
	basicConditionalSwitch()
	fmt.Printf("-----------------------------------\n\n")
}
