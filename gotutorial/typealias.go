package main

import "fmt"

//manytimes when we use constants we want to establish a custom type
//for the constant, thus we use the type keyword to create a new type
//that is an alias for an underlying type.

//Note go has 2 flavors of this, one is a true alias, the other is
//a new type that is based on the underlying type, but is not an alias
//for the underlying type.  This is a bit confusing, but it is a
//consequence of go's strong typing.

type myAppFlagType int //this is a NEW type that is based on int,
// but is not an alias
type myOtherAppFlagType = int //this is a just an ALIAS for int, so
//its syntactic sugar

// Now we can create constants of our custom types
const (
	//Lets create some constants based on our custom types
	FLAG_ONE myAppFlagType = iota
	FLAG_TWO
	FLAG_THREE

	FLAG_FOUR myOtherAppFlagType = iota
	FLAG_FIVE
	FLAG_SIX
)

func typeAliasTest() {
	fmt.Printf("What type is FLAG_TWO: %T\n", FLAG_TWO)
	fmt.Printf("What type is FLAG_FIVE: %T\n", FLAG_FIVE)
}

func typeAliasTest2() {
	t1 := FLAG_THREE
	t2 := FLAG_SIX

	//Also a mini lesson on type assertions...
	//create an interface that is the type of the constant
	//then use a type assertion to see if the interface is
	//the type of the constant
	f, ok := interface{}(t1).(int)
	f1, ok1 := interface{}(t2).(int)

	fmt.Printf("Is t1 a real int: %v, if so what is its value %v\n", ok, f)
	fmt.Printf("Is t2 a real int: %v, if so what is its value %v\n", ok1, f1)
}

func RunTypeAliasDemo() {
	fmt.Println("------ Running TypeAlias Demo ------")
	typeAliasTest()
	typeAliasTest2()
	fmt.Printf("-----------------------------------\n\n")
}
