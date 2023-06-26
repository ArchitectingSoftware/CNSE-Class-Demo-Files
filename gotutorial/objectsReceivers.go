package main

import "fmt"

//Now that we have seen basic structs, lets look how we can add behavior
//to structs.  This is a form of object oriented programming in go.
//The behavior is added by defining methods on the struct.

// define some base types
type car struct {
	make  string
	model string
	year  uint
	color string
}

func reviewTest() {
	//we have seen this before, we can create a struct by using the name
	//of the struct and then assigning the values of the fields.

	c := car{"Ford", "Mustang", 2019, "Red"}
	fmt.Println("CAR REVIEW: ", c)
}

// Note that many OO languages have a concept of constructors, go does not.
// Instead, you can create a function that returns a struct instance.
// Most of the time you return a pointer to the struct, but you can also
// return the struct itself.  The difference is that if you return the
// struct itself, a copy of the struct is made and returned.  If you
// return a pointer to the struct, then the struct is not copied and
// the caller gets a pointer to the struct.
func NewCar(make string, model string, year uint, color string) *car {
	c := &car{make, model, year, color}
	return c
}

//Now lets add some behavior to our struct.  We do this by defining
//methods on the struct.

// Go calls these receivers.  A receiver is defined right after the
// func keyword.
func (c *car) paint(color string) {
	c.color = color
}

func (c *car) getMake() string {
	return c.make
}

// Note that below we are not using a pointer receiver, this
// basically copies the entire struct providing it to the getModel()
// method.  This is not a good practice in general, but its good
// to know you can do this.
func (c car) getModel() string {
	return c.model
}

// Note that this is an example of why you should use pointer receivers
// unless there is a good reason not to.  In this case, we are changing
// the value of the model field, but since we are not using a pointer
// receiver, the change is not reflected in the original struct.
func (c car) setModelQuestionable(newModel string) {
	c.model = newModel
}

func RunReceiverTest() {

	c1 := NewCar("Ford", "Mustang", 2019, "Red")
	c2 := NewCar("Chevy", "Corvette", 2019, "Blue")

	fmt.Printf("Car1: %v\n", c1)
	fmt.Printf("Car2: %v\n", c2)

	//Lets paint the cars, this shows how we can update values
	//via pointer receivers
	c1.paint("Purple")
	c2.paint("Black")
	fmt.Printf("Painted Car1: %v\n", c1)
	fmt.Printf("Painted Car2: %v\n", c2)

	//Lets see how that non pointer receiver works, and how
	//many times its not what you want
	fmt.Printf("Car1 - PRE questionable receiver: %v\n", c1)
	c1.setModelQuestionable("Pinto")
	fmt.Printf("Car1 - POST questionable receiver - did model change?: %v\n", c1)
	fmt.Printf("-----------------------------------\n\n")
}

func RunObjectReceiverDemo() {
	fmt.Println("------ Running Object/Receiver Demo ------")
	reviewTest()
	RunReceiverTest()
	fmt.Printf("-----------------------------------\n\n")
}
