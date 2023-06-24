package main

import "fmt"

//In go structs are collections of fields.  They are useful for grouping data
//together to form records.  On the surface they seem very similar to structs in
//C but they can do a lot more.  We will see how we can add "methods" to structs
//to make them more powerful.  This can be used for object oriented programming in go

type person struct {
	name string
	age  int
}

// You can create a new struct by using the name of the struct and then assigning the
// values of the fields.  Note that the order of the fields does not matter, you can
// specify them in any order.
func createStructTest() {

	//Note in this case, Go assumes that the order of the values provided matches the
	//order of the fields in the struct.  This is not a good practice, but its good to
	//know
	p1 := person{"Bob", 20}
	fmt.Println(p1)

	//This is a bit better, we are explicitly specifying the field names, so the order
	//of the fields does not matter
	p2 := person{name: "Alice", age: 30}
	fmt.Println(p2)

	//We can also create a struct with each value initialized to its zero value. Zero
	//values in go are defaults. For example, the zero value for an int is 0, the zero
	//value for a string is "", the zero value for a bool is false, etc.
	p3 := person{}
	//We can then access the fields of the struct using the dot operator
	p3.name = "Fred"
	p3.age = 40
	fmt.Println(p3)
}

// You can also create a pointer to a struct.
// A couple of things to note here: In this example there is no real difference
// between this and the previous example.  The only difference is that we are
// using the & operator to indicate that we want a struct pointer. Go allocates
// memory for the structure on your behalf and garbage collects when the structure
// is no longer used.  This is different from C where you have to explicitly allocate
// and free memory for the structure.
func createStructPointerTest() {

	//Note in this case, Go assumes that the order of the values provided matches the
	//order of the fields in the struct.  This is not a good practice, but its good to
	//know
	p1 := &person{"Bob", 20}
	fmt.Println(p1)

	//This is a bit better, we are explicitly specifying the field names, so the order
	//of the fields does not matter
	p2 := &person{name: "Alice", age: 30}
	fmt.Println(p2)

	//We can also create a struct with each value initialized to its zero value. Zero
	//values in go are defaults. For example, the zero value for an int is 0, the zero
	//value for a string is "", the zero value for a bool is false, etc.
	p3 := &person{}
	//We can then access the fields of the struct using the dot operator
	p3.name = "Fred"
	p3.age = 40
	fmt.Println(p3)
}

func structAndPointerDemo() {

	//In this case p1 is created on the stack and initialized
	p1 := person{name: "Alice", age: 30}
	fmt.Println(p1)

	//In this case p2 is a copy of p1.  So changing p2 does not change p1
	p2 := p1
	p2.name = "Bob"
	p2.age = p1.age + 10
	fmt.Println(p1)
	fmt.Println(p2)

	p3 := person{name: "Joe", age: 50}
	fmt.Println(p3)
	p4 := &p3
	p4.name = "Fred"
	p4.age = p3.age + 10
	fmt.Println(p3)
	fmt.Println(p4)
}

//Lets see how this can also be used with functions

// Go does not have official constructors this is the idiomatic way to create a new
// person struct.  Notice the type is *person, this means that the function returns
// a pointer to a person struct.
func newPerson(name string) *person {
	//You can safely return a pointer to a local variable as a local variable
	//will survive the scope of the function
	return &person{name: name}
}

func newPersonTest() {
	p := newPerson("Jon")
	fmt.Println(p)
}

func newPersonValueCopy(name string) person {
	return person{name: name}
}

func newPersonValueCopyTest() {
	//In this case the person is copied back, this could get inefficent for large structs
	p := newPersonValueCopy("Jon")
	fmt.Println(p)
}

func RunStructsDemo() {
	fmt.Println("------ Running Structs Demo ------")

	createStructTest()
	createStructPointerTest()
	structAndPointerDemo()
	newPersonTest()
	newPersonValueCopyTest()
	fmt.Printf("-----------------------------------\n\n")
}
