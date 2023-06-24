package main

import "fmt"

//we have seen functions many times so far, but we have not yet looked
//at how to pass parameters and how to establish return types.  Note go
//allows you to return multiple values from a function, which is very
//useful in many cases.

//Idiomatic go uses this approach for functions to return a value and an error
//if one exists.  We will look at error handling in a later example.

// This function takes two ints and returns an int
func plus(a int, b int) int {
	//Go requires that you explicitly return from a function, it will not
	//automatically return the value of the last expression
	return a + b
}

// When you have multiple parameters of the same type, you can omit the type
// name for the like-typed parameters up to the final parameter that declares
// the type.  In this example, we have two parameters of type int, so we can
// omit the type name for the first parameter.
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Go supports multiple return values from a function.  This is very useful
// in many cases.  In this example, we return the sum and the difference of
// the two parameters.
func plusMinus(a, b int) (int, int) {
	return a + b, a - b
}

// Go supports named return values.  This is useful if you have multiple
// return values and you want to document what they are.  In this example,
// we return the sum and difference of the two parameters.  Note that we
// don't have to explicitly return the values, they are automatically
// returned.
func plusMinusNamed(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

// Go supports variadic functions, which can take a variable number of
// arguments.  In this example, we take an arbitrary number of ints and
// return the sum of them.
func plusMany(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Go supports closures, which are functions that reference variables
// defined outside of the function.  In this example, we define a function
// that returns a function that adds the parameter to the value defined
// outside of the function.
func plusX(x int) func(int) int {
	return func(y int) int { return x + y }
}

// Go supports recursion, which is when a function calls itself.  In this
// example, we define a function that calls itself until the parameter is
// equal to 0.
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Go supports anonymous functions, which are functions that are defined
// without a name.  In this example, we define an anonymous function and
// then call it.
func anonymousFunction() {
	func(msg string) {
		println(msg)
	}("hello")
}

// Go supports function values, which are functions that can be assigned
// to variables or passed as parameters.  In this example, we define a
// function value and then call it.
func functionValue() {
	f := func(msg string) {
		println(msg)
	}
	f("hello")
}

// Go supports function closures, which are functions that reference
// variables defined outside of the function.  In this example, we define
// a function that returns a function that adds the parameter to the value
// defined outside of the function.
func functionClosure() {
	f := plusX(10)
	println(f(5))
}

func RunFunctionsDemo() {
	fmt.Println("------ Running Functions Demo ------")
	fmt.Printf("1 + 2 = %d\n", plus(1, 2))
	fmt.Printf("1 + 2 + 3 = %d\n", plusPlus(1, 2, 3))
	sum, diff := plusMinus(1, 2)
	fmt.Printf("1 + 2 = %d, 1 - 2 = %d\n", sum, diff)
	sum, diff = plusMinusNamed(1, 2)
	fmt.Printf("1 + 2 = %d, 1 - 2 = %d\n", sum, diff)
	fmt.Printf("1 + 2 + 3 + 4 + 5 = %d\n", plusMany(1, 2, 3, 4, 5))
	fmt.Printf("10 + 5 = %d\n", plusX(10)(5))
	fmt.Printf("5! = %d\n", factorial(5))
	anonymousFunction()
	functionValue()
	functionClosure()
	fmt.Printf("-----------------------------------\n\n")
}
