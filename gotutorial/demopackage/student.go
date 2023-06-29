package demopackage

import "fmt"

//This is a demonstration of how to use packages in go that manage
//visibility via controlling how things are exported and kept
//private in go.
//Note that the package name at the top is demopackage, by convention
//we name packages the same as the directory they are in.  This is
//not required, but it is a good practice.

type Student struct {
	Name           string
	Year           int
	GPA            float32
	ssn            string
	tuitionBalance float32
}

var initialBalance = 1000.0 //this is private, you cannot see it
//outside of the package

const PackageName = "demopackage" //this is public, you can see it
const author = "Brian Mitchell"   //this is private, you cannot see it

func New() *Student {
	newStudent := &Student{
		Name:           "John Doe",
		Year:           1,
		GPA:            4.0,
		ssn:            "123-45-6789",
		tuitionBalance: float32(initialBalance),
	}
	//Note by default go uses float64 for floating point, if you want
	//to use float32 you need to explicitly specify it.  This is the
	//way go does type conversion. The syntax is different than C

	fmt.Println("package:", PackageName, "author:", author)

	return newStudent
}

func (s *Student) AdjustBalance(amount float32) float32 {
	s.tuitionBalance += amount
	return s.tuitionBalance
}

func (s *Student) GetSSN(password string) string {
	if password == s.getPassword() {
		return s.ssn
	}
	return "NOT_AUTHORIZED"
}

func (s *Student) getPassword() string {
	return "secret"
}
