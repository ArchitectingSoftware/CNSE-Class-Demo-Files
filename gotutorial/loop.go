package main

import "fmt"

//Go really only has one looping construct, the for loop.  It can be used in
//multiple ways, but it is the only looping construct in the language.

func basicLoopTest() {
	//Go's basic for loop has three components separated by semicolons:
	//Its very similar to C's for loop.  Note that the variable i
	//is only scoped to the for loop.
	for i := 0; i < 10; i++ {
		println(i)
	}
}

func basicLoopTest1() {
	//You can also use the break statment inside of a loop to exit it
	for {
		fmt.Println("loop")
		break
	}
}

func basicLoopTest2() {
	//You can also use a for loop like a while loop, notice any condition on
	//the for expression that evaluates to a boolean is valid.  If the condition
	//is true, the loop will continue, if it is false, the loop will break.
	i := 0
	for i < 10 {
		println(i)
		i++
	}
}

func rangeLoopTest() {
	//You can use the range keyword to loop over an array, slice, string, map, or channel
	//The range keyword returns the index and the value of the current element
	//Note that the index is optional, if you don't need it, you can use the _ character
	//to ignore it

	//In this first example, we are looping over an slice, notice that range returns 2
	//values, the first is the index, the second is the value.  We are ignoring the index
	//in this example.  In general, if you use the underscore in go, you are telling the
	//compiler that you don't care about that value.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//In this example, we are looping over a map.  Notice that this time we are using
	//both the key and the value
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}

func RunLoopDemo() {
	fmt.Println("------ Running Loop Demo ------")
	basicLoopTest()
	basicLoopTest1()
	basicLoopTest2()
	rangeLoopTest()
	fmt.Printf("-----------------------------------\n\n")
}
