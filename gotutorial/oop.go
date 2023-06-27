package main

import "fmt"

// define some base types
type individual struct {
	name string
	age  uint
}

type student struct {
	individual
	gpa float64
}

// Now define some additional types that use embedding, this is a form
// of inheritance that is based on composition rather than inheritance
// as in traditional OOP languages.  This is a more flexible approach
// that allows for multiple inheritance and avoids the problems of
// the diamond problem.
type employee struct {
	title          string
	salary         float64
	privacyRequest bool
}

type professor struct {
	individual
	employee
	department string
	nameSuffix string
}

// Now lets add some behavior to our types
func (i *individual) getName() string {
	return i.name
}

func (i *individual) setName(name string) {
	i.name = name
}

func (p *professor) setNameSuffix(suffix string) {
	p.setName(p.getName() + " " + suffix)
}

func (i *individual) getAge() uint {
	return i.age
}

func (p *professor) getAge() uint {
	//This is kind of a made up example, but it shows how you can
	//override methods in embedded types.  In this case we are
	//checking to see if the privacyRequest flag is set, if so
	//we return 0, otherwise we return the age from the embedded
	//individual type.
	//
	//This is similar to using "super" in Java or C++.
	age := p.individual.getAge()
	if p.privacyRequest {
		return 0
	} else {
		return age
	}
}

func oopDemo1() {

	//Lets create a new professor, note we represent the
	//composed types as fields in the struct using the typename
	prof := professor{
		individual: individual{name: "John Doe",
			age: 42,
		},
		employee: employee{title: "Professor",
			salary:         100000,
			privacyRequest: true,
		},
		department: "Computer Science",
		nameSuffix: "PhD",
	}

	// Now lets call some methods on our professor.  Note this will
	//automatically call getName from individual
	fmt.Println("Prof getName()", prof.getName())

	// Now lets call getAge, note that this will call the overriden getAge
	// method in professor
	prof.setNameSuffix("PhD")
	println("Prof: getName() - post Suffix Set", prof.getName())

	// Now lets create a student
	s := student{
		individual: individual{
			name: "Jane Student",
			age:  20},
		gpa: 3.5,
	}

	//note that this will call the individual getName method because
	//student does not override it
	fmt.Println("Student getName()", s.getName())
}

func RunOopDemo() {
	fmt.Println("------ Running OOP Demo ------")
	oopDemo1()
	fmt.Printf("-----------------------------------\n\n")
}
