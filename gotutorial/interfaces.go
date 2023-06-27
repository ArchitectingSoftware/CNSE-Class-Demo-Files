package main

import (
	"fmt"
	"strconv"
)

//Go like many other languages has the concept of interfaces.  An interface
//is a contract that defines a set of methods that a type must implement
//in order to be considered a valid implementation of the interface.
//Interfaces are a powerful feature of Go, and are used extensively in
//the standard library.  Interfaces are also used to implement polymorphism
//in Go.

//What makes go distinctively different from other languages is that
//interfaces are implemented implicitly.  This means that a type does
//not have to explicitly declare that it implements an interface, it
//just has to implement the methods defined in the interface.  This
//is a very powerful feature of Go, and allows for a lot of flexibility
//in how interfaces are used.

//Lets look at a simple example of an interface.  We will define a
//simple interface that defines a method called "getName".  We will
//then define a type that implements this interface.  We will then
//demonstrate how to use the interface to call the method on the
//type that implements the interface.

type nameable interface {
	GetName() string
}

// Note the three types below, they dont have anything in common, but
// they are all things that are "nameable".
type university struct {
	name     string
	location string
}

type pet struct {
	name  string
	breed string
}

type boat struct {
	name            string
	motorHorsePower int
}

func (p *boat) GetName() string {
	return "Hi I am a boat named : " + p.name +
		" and I have a motor with " +
		strconv.Itoa(p.motorHorsePower) + " horsepower"
}

func (p *pet) GetName() string {
	return "Hi my owners named me: " + p.name + " and I am a " + p.breed
}

func (u *university) GetName() string {
	return "Note that " + u.name + " is located at " + u.location
}

func InterfaceDemoTest() {
	//Now lets create some instances of these types
	u := university{name: "Drexel", location: "Philadelphia"}
	p := pet{name: "Rover", breed: "Lab"}
	b := boat{name: "The Titanic", motorHorsePower: 1000000}

	//Now lets create a slice of nameable things
	things := []nameable{&u, &p, &b}

	//Now lets iterate over the slice and call the GetName method
	for _, thing := range things {
		println(thing.GetName())
	}
}

func RunInterfaceDemo() {
	fmt.Println("------ Running Interface Demo ------")
	InterfaceDemoTest()
	fmt.Printf("-----------------------------------\n\n")
}
