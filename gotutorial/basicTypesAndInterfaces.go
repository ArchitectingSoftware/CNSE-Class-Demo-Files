package main

import (
	"fmt"
)

func basicTypes() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	g := `In go you can use backticks to create
	a multi-line string. This is useful for
	creating SQL statements, JSON, or HTML.`
	fmt.Println(g)
}

func typeSizes() {
	var i1 int = 1
	var i2 int8 = 2
	var i3 int16 = 3
	var i4 int32 = 4
	var i5 int64 = 5

	//note in go you can use %v to print out the value of a variable regardless of type
	fmt.Printf("Integers: %v %v %v %v %v\n", i1, i2, i3, i4, i5)

	//note that you can specify type sizes for floats as well
	var f1 float32 = 1.0
	var f2 float64 = 2.0
	fmt.Printf("Floats: %v %v\n", f1, f2)

	//note that you can specify type sizes for unsigned ints as well
	var u1 uint = 1
	var u2 uint8 = 2
	var u3 uint16 = 3
	var u4 uint32 = 4
	var u5 uint64 = 5
	fmt.Printf("Unsigned Integers: %v %v %v %v %v\n", u1, u2, u3, u4, u5)
}

func typeConversions() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Printf("Conversions: %v %v %v\n", i, f, u)
}

func typeInference() {
	v := 42
	fmt.Printf("v is of type %T\n", v)

	v2 := 42.0
	fmt.Printf("v2 is of type %T\n", v2)

	//complex numbers are built in to go
	v3 := 42 + 0.5i
	fmt.Printf("v3 is of type %T\n", v3)

	v4 := "42"
	fmt.Printf("v4 is of type %T\n", v4)

	v5 := `42`
	fmt.Printf("v5 is of type %T\n", v5)

	v6 := `42 + 0.5i`
	fmt.Printf("v6 is of type %T\n", v6)
}

func interfaceDemo() {
	//Note in go, the empty interface is the interface that has no methods
	//it is satisfied by any type
	var s1 interface{} = "hello"
	var s2 interface{} = 1
	var s3 interface{} = true

	fmt.Printf("Empty Interface dealing with types: %v %v %v\n", s1, s2, s3)

	//After go 1.18, you can use the any keyword to do the same thing
	var a1 any = "hello"
	var a2 any = 1
	var a3 any = true

	fmt.Printf("Empty Interface dealing with types with any: %v %v %v\n", a1, a2, a3)
}

func typeAssertions() {
	//Type assertions are used to extract the underlying value of the interface
	//and check its type
	var i interface{} = "hello"

	//this is a type assertion
	s := i.(string)
	fmt.Println(s)

	//this is a type assertion with a second return value that indicates if the assertion succeeded
	s, ok := i.(string)
	fmt.Println(s, ok)

	//this is a type assertion that will fail
	f, ok := i.(float64)
	fmt.Println(f, ok)

	//this is a type assertion that will panic, and will end the program
	//we comment it out so that this does not happen, but give it a try yourself
	/*
		f = i.(float64)
		fmt.Println(f)
	*/
}

func RunTypesAndInterfacesDemo() {
	fmt.Println("------ Running Basic Types/Interfaces Demo ------")
	basicTypes()
	typeSizes()
	typeConversions()
	typeInference()
	interfaceDemo()
	typeAssertions()
	fmt.Printf("-----------------------------------\n\n")
}
