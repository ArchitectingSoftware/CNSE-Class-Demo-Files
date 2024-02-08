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

type asset interface {
	GetValue() float64
}

type nameable interface {
	GetName() string
}

// Note the three types below, they dont have anything in common, but
// they are all things that are "nameable".
type CashAccount struct {
	Value float64
}

func (c *CashAccount) AddAsset(a asset) {
	c.Value += a.GetValue()
}

type InvestmentAccount struct {
	Value float64
}

func (c *InvestmentAccount) AddAsset(a asset) {
	c.Value += a.GetValue()
}

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

func (p *boat) GetValue() float64 {
	return 100000.00
}

func (p *pet) GetName() string {
	return "Hi my owners named me: " + p.name + " and I am a " + p.breed
}

func (u *pet) GetValue() float64 {
	return 100.00
}

func (u *university) GetName() string {
	return "Note that " + u.name + " is located at " + u.location
}

func (u *university) GetValue() float64 {
	return 1000000000.00
}

func InterfaceDemoTest() {
	//Now lets create some instances of these types
	u := university{name: "Drexel", location: "Philadelphia"}
	p := pet{name: "Rover", breed: "Lab"}
	b := boat{name: "The Titanic", motorHorsePower: 1000000}

	fmt.Println("Univ: ", u.GetName())
	fmt.Println("Pet: ", p.GetName())
	fmt.Println("Boat: ", b.GetName())

	var n1, n2, n3 nameable
	n1 = &u
	n2 = &p
	n3 = &b

	println(n1.GetName())
	println(n2.GetName())
	println(n3.GetName())

	var a1, a2 asset
	a1 = &u
	a2 = &b

	println(a1.GetValue())
	println(a2.GetValue())

	//Now lets create a slice of nameable things
	things := []nameable{&u, &p, &b}

	//Now lets iterate over the slice and call the GetName method
	for _, thing := range things {
		println(thing.GetName())
	}

	invAcct := InvestmentAccount{Value: 1000.0}
	cashAcct := CashAccount{Value: 100.0}

	invAcct.AddAsset(&u)
	cashAcct.AddAsset(&b)
	cashAcct.AddAsset(&p)

	fmt.Println("Investment Account Value: ", invAcct.Value)
	fmt.Println("Cash Account Value: ", cashAcct.Value)
}

func RunInterfaceDemo() {
	fmt.Println("------ Running Interface Demo ------")
	InterfaceDemoTest()
	fmt.Printf("-----------------------------------\n\n")
}
