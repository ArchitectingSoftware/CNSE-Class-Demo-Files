package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//this demo will show various ways to do IO in Go

//The code below is from GoByExample.com

// helper since fileIO requires a lot of error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFileDemo() {

	//Perhaps the most basic file reading task is slurping a file's
	//entire contents into memory.
	dat, err := os.ReadFile("./tmp/dat")
	check(err)
	fmt.Print(string(dat))

	//You'll often want more control over how and what parts of a file
	//are read. For these tasks, start by Opening a file to obtain an
	//os.File value.
	f, err := os.Open("./tmp/dat")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	//Seeking to a known position
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	//No reset, just seek to 0,0
	_, err = f.Seek(0, 0)
	check(err)

	//Similiar to java, you can use a buffered reader to
	//improve performance
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}

func writeFileDemo() {

	//Writing files in Go follows similar patterns to the
	//ones we saw earlier for reading.
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("./tmp/dat1", d1, 0644)
	check(err)

	//For more granular writes, open a file for writing.
	f, err := os.Create("./tmp/dat2")
	check(err)

	//It's idiomatic to defer a Close immediately after opening a file.
	//This will execute at the end of the enclosing function (main),
	//after writeFileDemo has finished.  This prevents resource leaks
	//that you see in other languages
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	//bufio provides buffered writers in addition to the buffered
	//readers we saw earlier.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}

func RunFileIODemo() {
	fmt.Println("------ Running FileIO Demo ------")
	readFileDemo()
	writeFileDemo()
	fmt.Printf("-----------------------------------\n\n")
}
