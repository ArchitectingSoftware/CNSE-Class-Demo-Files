package main

import "fmt"

func aboutArrays() {
	//Go has arrays
	//Note that arays are somewhat inflexible in that the size is allocated at compile time
	//and cannot be changed.  We will examine slices next which are more flexible.

	var x [5]int //this is how you create an array, notice the size is part of the type
	x[4] = 100
	println(x[4])

	l := len(x) //this is how you get the length of an array, len is a built in function
	println(l)
}

func aboutSlices() {
	//Go has slices
	//Note that slices are more flexible than arrays in that the size is allocated at runtime
	//and can be changed.

	var x []int      //this is how you create a slice, notice the size is not part of the type
	x = append(x, 1) //append is a built in function
	println(x[0])

	l := len(x) //this is how you get the length of a slice, len is a built in function
	println(l)
}

func initializingSlices() {
	//Go has a shorthand way of initializing slices
	x := []int{1, 2, 3, 4, 5}
	fmt.Printf("x: %v\n", x)
}

func preallocatedSlices() {
	//Go has a make function that can be used to preallocate slices
	//This is useful if you know the size of the slice ahead of time
	//or if you want to allocate a slice with a specific capacity

	x := make([]int, 5) //this makes a slice with an initial size of 5
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5
	fmt.Printf("x: %v\n", x)
	//note that x[5] = 5 will cause an error because the slice is only size 5
}

func growingSlices() {
	//Go can grow slices by using the append function
	//Note that the append function returns a new slice, it does not modify the
	//original slice.  Go tries to be efficient about this and will double the
	//size of the slice when it grows.

	x := make([]int, 5) //this makes a slice with an initial size of 5
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5
	fmt.Printf("x (before growth): %v\n", x)
	x = append(x, 6)           //adding an element to a slice
	x = append(x, 7, 8, 9, 10) //note append can take multiple arguments
	fmt.Printf("x (after growth): %v\n", x)
}

func indexingSlices() {
	//Go can index slices and return parts of them

	x := []int{1, 2, 3, 4, 5}

	y := x[1:3] //this is how you get a slice of a slice
	//notice how the indexing works, this is called a half open range
	//the first index is inclusive and the second index is exclusive
	//so this will get elements 1 and 2.
	fmt.Printf("x: %v\n", y)

	//You can also omit the first index and go will assume 0
	//This will get elements 0 and 1
	z := x[:2] //this will get elements 0 and 1
	fmt.Printf("z: %v\n", z)

	//You can also omit the second index and go will assume the length of the slice
	a := x[2:] //this will get elements 2, 3, and 4
	fmt.Printf("a: %v\n", a)
}

func aboutMaps() {
	//Go has maps
	//Maps are a key value store.  Go manages all of the storage
	//associated with maps

	var x map[string]int //note that this is a variable that is a map type, its
	//value is nil until you make it a map by using the make
	// function

	x = make(map[string]int) //this makes a map
	x["key"] = 10
	x["key2"] = 20
	fmt.Printf("x: %v\n", x)

	l := len(x) //this is how you get the length of a map, len is a built in function
	println(l)

	//of course you can also declare and initialize a map in one line
	y := map[string]int{"key1": 10, "key2": 20}
	fmt.Printf("y: %v\n", y)

	//you can also make a map with make and a capacity, this might improve memory
	//usage and efficiency if you know the size ahead of time.  Go will grow the
	//map dynamically if you need more space
	z := make(map[string]string, 100)
	z["key"] = "value"
	fmt.Printf("z: %v\n", z)
}

func loopingOverMapsDemo() {
	x := make(map[string]int) //this makes a map
	x["key"] = 10
	x["key2"] = 20
	x["key3"] = 30
	x["key4"] = 40

	//you can loop over a map using the range keyword
	//note that the order of the keys is not guaranteed
	for key, value := range x {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

	//you can also just get the keys
	for key := range x {
		fmt.Printf("key: %v\n", key)
	}

	//you can also just get the values
	for _, value := range x {
		fmt.Printf("value: %v\n", value)
	}
}

func deleteFromMapDemo() {
	x := make(map[string]int) //this makes a map
	x["key"] = 10
	x["key2"] = 20
	x["key3"] = 30
	x["key4"] = 40

	//lets say we want to remove key2 from the map

	//we should first make sure the key exists
	if _, ok := x["key2"]; ok {
		fmt.Printf("key2 exists in the map\n")

		//now we can delete it
		delete(x, "key2")
		fmt.Printf("x: %v\n", x)
	} else {
		fmt.Printf("key2 does not exist in the map\n")

		//note that delete will not cause any harm if a key
		//does not exist in the map, its basicaly a noop
		//but its better to check first
		delete(x, "key2")
	}

}

func RunArraySliceMapDemo() {
	fmt.Println("------ Running Array/Slice/Map Demo ------")
	aboutArrays()
	aboutSlices()
	initializingSlices()
	preallocatedSlices()
	growingSlices()
	indexingSlices()
	aboutMaps()
	loopingOverMapsDemo()
	deleteFromMapDemo()
	fmt.Printf("-----------------------------------\n\n")
}
