package main

import "fmt"

// You can create a constant block that auto assigns values using the iota keyword
// Note these are global constants, the first letter is capitalized
const (
	FIRST_OPTION = iota
	SECOND_OPTION
	THIRD_OPTION
)

const (
	INITAL_VALUE_OPTION = iota + 5
	SECOND_VALUE_OPTION
	THIRD_VALUE_OPTION
)

// You can also create a constant block that assigns values individually
// Note these are not global constants, the first letter is lower case
const (
	localConstant    = "local"
	localConstNumber = 20
	localConstBool   = true
	localConstFloat  = 3.14
)

func basicConstantsTest() {
	//Go has constants
	const x string = "Hello, World"
	println(x)
}

func basicConstantsTest2() {
	fmt.Printf("Constants2: %v %v %v\n", FIRST_OPTION, SECOND_OPTION, THIRD_OPTION)
}

func basicConstantsTest3() {
	fmt.Printf("Constants3: %v %v %v\n", INITAL_VALUE_OPTION, SECOND_VALUE_OPTION, THIRD_VALUE_OPTION)
}

func basicConstantsTest4() {
	fmt.Printf("Constants4: %v %v %v %v\n", localConstant, localConstNumber, localConstBool, localConstFloat)
}

func RunConstantsDemo() {
	fmt.Println("------ Running Constants Demo ------")
	basicConstantsTest()
	basicConstantsTest2()
	basicConstantsTest3()
	basicConstantsTest4()
	fmt.Printf("-----------------------------------\n\n")
}
