package main

import "fmt"

//Note go has pointers, but no pointer arithmetic.  You can use the & operator
//to get the address of a variable and the * operator to dereference a pointer.

func basicPointersTest() {
	//We can use the & operator to get the address of a variable
	i := 1
	println("initial:", i)
	println("pointer:", &i)

	//We can use the * operator to dereference a pointer
	p := &i
	println("pointer:", p)
	println("dereferenced:", *p)

	//We can also change the value of the variable through the pointer
	*p = 2
	println("changed:", i)
}

// With pointers you can pass by reference, which means that the function
// can change the value of the variable passed in.  This is useful if you
// want to change the value of the variable passed in.
func passByReferenceTest() {
	//In this example, we pass in a pointer to an int and change the value
	//of the variable through the pointer.
	i := 1
	println("initial:", i)
	passByReference(&i)
	println("changed:", i)
}

func passByReference(i *int) {
	*i = 2
}

// With pointers you can pass by value, which means that the function
// cannot change the value of the variable passed in.  This is useful if you
// don't want the function to change the value of the variable passed in.
func passByValueTest() {
	//In this example, we pass in an int and try to change the value
	//of the variable.  Notice that the value of the variable is not changed.
	i := 1
	println("initial:", i)
	passByValue(i)
	println("changed:", i)
}

func passByValue(i int) {
	i = 2
}

//At the end of the day, pass by value means that the variable is COPIED
//so any operations on the variable will not affect the original variable.
//Pass by reference means that the variable is NOT COPIED, its memory address
//is passed so any operations on the variable will affect the original variable.

//Note that with simple types like int, bool the copying operation is not a big
//deal because these dont take a lot of memory.  However, if you have a large
//struct, copying it can be very expensive, so you should pass by reference in
//that case.  We will look at structs soon.

func RunBasicPointersDemo() {
	fmt.Println("------ Running Basic Pointers Demo ------")
	basicPointersTest()
	passByReferenceTest()
	passByValueTest()
	fmt.Printf("-----------------------------------\n\n")
}
