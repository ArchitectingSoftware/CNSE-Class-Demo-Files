package oopobjects

type individual struct {
	Name       string
	age        uint
	agePrivacy bool
}

type Student struct {
	individual
	GPA float64
}

type Employee struct {
	individual
	Title  string
	salary float64
	ssn    string
}

type Professor struct {
	Employee
	Department string
	NameSuffix string
}

// Now lets add some behavior to our types
func (i *individual) GetAge() uint {
	return i.age
}

func (p *Professor) SetNameSuffix(suffix string) {
	p.Name = p.Name + ", " + suffix
}

func (p *Professor) GetAge() uint {
	//This is kind of a made up example, but it shows how you can
	//override methods in embedded types.  In this case we are
	//checking to see if the privacyRequest flag is set, if so
	//we return 0, otherwise we return the age from the embedded
	//individual type.
	//
	//This is similar to using "super" in Java or C++.
	age := p.individual.GetAge()
	if p.agePrivacy {
		return 0
	} else {
		return age
	}
}

var testStudent = Student{
	individual: individual{
		Name:       "Jane Student",
		agePrivacy: false,
		age:        20,
	},
	GPA: 3.5,
}

// Note if we statically initialize a struct that embeds other structs we
// must initialize the embeded structs by type.  See the NewProfFromScratch
// function below for an example of how we can directly initialize a struct
// using . notation to reference the embedded types indirectly.
var testProfessor = Professor{
	Employee: Employee{
		individual: individual{
			Name:       "John Doe",
			agePrivacy: false,
			age:        42,
		},
		Title:  "Professor",
		salary: 100000,
		ssn:    "123-45-6789",
	},
	Department: "Computer Science",
	NameSuffix: "PhD",
}

var testEmployee = Employee{
	individual: individual{
		Name:       "Tim Electrician",
		agePrivacy: true,
		age:        50,
	},
	Title:  "Electrician",
	salary: 80000,
	ssn:    "987-65-4321",
}

func (e *Employee) GiveRaise(percent float64) (float64, float64) {
	oldSalary := e.salary
	e.salary = e.salary * (1 + percent)
	return oldSalary, e.salary
}

func (e *Employee) Hire() *Employee {
	return &testEmployee
}

func (p *Professor) Hire() *Professor {
	return &testProfessor
}

func (s *Student) Enroll() *Student {
	return &testStudent
}

func GetProfessor() *Professor {
	return &testProfessor
}

func GetEmployee() *Employee {
	return &testEmployee
}

func GetStudent() *Student {
	return &testStudent
}

func NewProfFromScratch() *Professor {
	//notice from above if we statically initialize a professor we need to
	//initialize the embedded types as well.  This is a bit of a pain.
	//If we use . notation we can reference all the fields of the embedded
	//types directly
	p := &Professor{}
	p.Name = "John Doe"
	p.age = 42
	p.agePrivacy = false
	p.Title = "Professor"
	p.salary = 100000
	p.ssn = "123-45-6789"
	p.Department = "Computer Science"
	p.NameSuffix = "PhD"
	return p
}
