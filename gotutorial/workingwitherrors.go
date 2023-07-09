package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

//Go is a lot different than many other languages when it comes to
//dealing with errors.  Go does not have the concept of a try/catch
//block.  Errors are expected to be handled by the caller, within
//the context of making calls that can generate errors.  Some people
//think this is a killer feature of the language, others say that it
//causes a lot of boilerplate code.

// Lets look at an example of how to handle errors in go.  We will
// use the strconv package to convert a string to an int.  This
// function can return an error if the string cannot be converted
// to an int.
func errorsTest() {
	//lets try to convert a string to an int
	//the function we are calling returns 2 values, the int and an error
	//we use the _ to ignore the int value
	_, err := strconv.Atoi("1234")

	//notice the error is handled right after the call, this is what
	//you will see in most go code.
	if err != nil {
		fmt.Printf("Error converting string to int: %v\n", err)
	}
}

func IdomaticErrorFuncDemo(word string) (string, error) {
	//This is a stupid example but shows the way many functions
	//in go are written, they return the desired value, and and
	//error.  The error is nil if there is no error, otherwise
	//it is populated with error information.

	//Many other languages override a standard return type to
	//indicate an error.  For example in C, the socket functions
	//return useful data as an int if successful, or a negative
	//value to indicate an error.  Java is famous for using null
	//to indicate an error, so the return value is valid if not
	//an error, and null indicates an error has occoured.

	if len(word) > 5 {
		return word, nil
	} else {
		return "", errors.New("Word is too long")
	}
}

func idomaticErrorFuncDemo2() {
	s, err := IdomaticErrorFuncDemo("The")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Success: %v\n", s)
	}

	s, err = IdomaticErrorFuncDemo("The quick brown fox jumped over the lazy dog")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Success: %v\n", s)
	}

	//sometimes we see shortcuts where the erorr is handled within
	//the context of the function call itself, lets look at this
	if s, err := IdomaticErrorFuncDemo("Hello class"); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Success: %v\n", s)
	}
}

//Many times an error is just a string, but sometimes it is
//useful to create a custom type for an error.  This allows
//you to create a set of errors that are related, and then
//you can check for specific errors in your code.  The error
//interface is very simple, it is just a method that returns
//a string.  So if you create a custom type that implements
//this method, you can use it as an error.
//
// type error interface {
//   Error() string
// }

// Here is one way to do it, we create a custom type that
// implements the error interface.
type StringTooLongError struct{}

func (m *StringTooLongError) Error() string {
	return "The String is Too Long"
}

// if we just want to create a custom error, one way to do it is
// to use the error.New() function/constructor
var StringTooShortError = errors.New("The string is too short")

// You can also use the fmt.Errorf() function to create a custom
// error. This is useful if you want to take advantage of formatting
var StringTooLongFormattedError = fmt.Errorf("The provided string is %s", "too long")

func idomaticErrorFuncDemo3(msg string) error {
	_, err := IdomaticErrorFuncDemo(msg)
	if err != nil {
		//Becuase this is a type, we need to create an error instance then we can return it
		myErr := &StringTooLongError{}
		return myErr
	}

	return nil
}

func idomaticErrorFuncDemo4(msg string) error {
	if len(msg) < 5 {
		return StringTooShortError
	}

	if len(msg) > 10 {
		return StringTooLongFormattedError
	}

	return nil
}

// Since errors are just string wrappers, it sometimes is useful to create a more useful
// error type.  For example, in web development, a status code is often useful.  Lets look
// at this.
type NetworkError struct {
	StatusCode int
	Msg        string
}

var NetError = &NetworkError{} //This is a global variable that is a pointer to a NetworkError
//We use this technique if we want to look for a specific error
//type in our code.

func NewNetworkError(statusCode int, msg string) *NetworkError {
	return &NetworkError{statusCode, msg}
}
func (e *NetworkError) Error() string {
	return fmt.Sprintf("Network Error: %v, %v", e.StatusCode, e.Msg)
}

func SimulateNetworkCall() error {
	r := rand.Intn(2) //Generates a random number between 0 and 1

	if r == 0 {
		return NewNetworkError(500, "Internal Server Error")
	}

	return nil
}

func SimulateNetworkCall2Helper() error {
	r := rand.Intn(3) //Generates a random number between 0 and 1

	fmt.Printf("Random Number: %v\n", r)

	switch r {
	case 0:
		return NewNetworkError(500, "Internal Server Error")
	case 1:
		return errors.New("Nothing really happened of interest")
	default:
		return nil
	}
}

func SimulateNetworkCall2() error {
	err := SimulateNetworkCall2Helper()

	if errors.Is(err, NetError) {
		fmt.Printf("Network Error: %v\n", err)
	}
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("No Error")
	return err
}

func SimulateNetworkCall3() error {
	err := SimulateNetworkCall2Helper()

	_, netError := err.(*NetworkError)
	if netError {
		fmt.Printf("Network Error: %v\n", err)
	}
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("No Error")
	return err
}

func RunErrorsDemo() {
	fmt.Println("------ Running Errors Demo ------")
	errorsTest()
	idomaticErrorFuncDemo2()
	idomaticErrorFuncDemo3("Hello")
	idomaticErrorFuncDemo3("Hello Class")
	idomaticErrorFuncDemo4("Hello")
	idomaticErrorFuncDemo4("Hello Class")
	SimulateNetworkCall()
	SimulateNetworkCall2()
	SimulateNetworkCall3()
	fmt.Printf("-----------------------------------\n\n")
}
