package main

import (
	"encoding/json"
	"fmt"
)

//JSON is a common format used for data exchange, especially in cloud
//native applications.  Go has excellent support for JSON encoding and
//decoding, built-in.

// Lets start with a basic structure for our data.  Notice that all of
// the fields start with capital letters.  THIS IS IMPORTANT, as the
// built in JSON encoding will only encode "exported" fields, remember
// Go has the convention that capitialized fields are exported, and
// lowercase fields are not.
type ToDoItem struct {
	Id      int
	Title   string
	Details string
	DueDate string
	IsDone  bool
}

// Lets create a slice of ToDoItems to work with
var ToDoList []ToDoItem

// Lets create a function to populate our slice
func populateToDoList() {
	ToDoList = append(ToDoList, ToDoItem{Id: 1, Title: "Get Milk", Details: "2%", DueDate: "2021-01-01", IsDone: false})
	ToDoList = append(ToDoList, ToDoItem{Id: 2, Title: "Get Bread", Details: "Whole Wheat", DueDate: "2021-01-01", IsDone: false})
	ToDoList = append(ToDoList, ToDoItem{Id: 3, Title: "Get Eggs", Details: "Large", DueDate: "2021-01-01", IsDone: false})
	ToDoList = append(ToDoList, ToDoItem{Id: 4, Title: "Get Beer", Details: "IPA", DueDate: "2021-01-01", IsDone: false})
}

func testToDoToJson() {
	//Lets populate our slice
	populateToDoList()

	//Lets convert our slice to JSON
	//Note that the second parameter is an error, so we need to check
	//it
	jsonBytes, err := json.Marshal(ToDoList)
	if err != nil {
		fmt.Println(err)
	}

	//Lets print the JSON
	fmt.Println("This will just print the string in JSON")
	fmt.Println(string(jsonBytes))
}

func testToDoToJsonPrettyPrint() {
	//Lets populate our slice
	populateToDoList()

	//Lets convert our slice to JSON
	//Note that the second parameter is an error, so we need to check
	//it
	jsonBytes, err := json.MarshalIndent(ToDoList, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	//Lets print the JSON
	fmt.Println("Lets print formatted...")
	fmt.Println(string(jsonBytes))
}

// Note that by default the JSON encoder will use the name of the field
// as the JSON key.  Go structurs have a special annotation that can be
// used to override this behavior.  They are called "tags" and are
// placed after the field name, and are surrounded by backticks. Below
// we are renaming all of our fields

// also note we can add other information to the tag, such as omitempty
// which will cause the encoder to skip the field if it is empty.  By
// default go will encode all fields, even if they are empty (using their
// zero value)
type ToDoItem2 struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
	DueDate string `json:"due,omitempty"`
	IsDone  bool   `json:"finished"`
}

// Lets create a slice of ToDoItems to work with
var ToDoList2 []ToDoItem2

// Lets create a function to populate our slice
func populateToDoList2() {
	ToDoList2 = append(ToDoList2, ToDoItem2{Id: 1, Title: "Get Milk", Details: "2%", DueDate: "2021-01-01", IsDone: false})
	ToDoList2 = append(ToDoList2, ToDoItem2{Id: 2, Title: "Get Bread", Details: "Whole Wheat", DueDate: "2021-01-01", IsDone: false})
	ToDoList2 = append(ToDoList2, ToDoItem2{Id: 3, Title: "Get Eggs", Details: "Large", IsDone: false})
	ToDoList2 = append(ToDoList2, ToDoItem2{Id: 4, Title: "Get Beer", Details: "IPA", IsDone: false})
}

func testAnnotatedJson() {
	//Lets populate our slice
	populateToDoList2()

	//Lets convert our slice to JSON
	//Note that the second parameter is an error, so we need to check
	//it
	jsonBytes, err := json.MarshalIndent(ToDoList2, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	//Lets print the JSON
	fmt.Println("Lets print formatted with our custom tag names...")
	fmt.Println(string(jsonBytes))
}

func simpleJsonString() {
	ToDoListItem := ToDoItem2{
		Id:      4,
		Title:   "Get Beer",
		Details: "IPA",
		IsDone:  false}

	//Lets convert our structure to a JSON string
	jsonBytes, err := json.MarshalIndent(ToDoListItem, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	jsonString := string(jsonBytes)

	//Lets print the JSON
	fmt.Println("Lets see that we have a starting string")
	fmt.Println(jsonString)

	//Now lets convert it back to a struct
	var myToDoItem ToDoItem2
	err = json.Unmarshal([]byte(jsonString), &myToDoItem)
	if err != nil {
		fmt.Println(err)
	}
	//Lets print the struct and make sure all is OK
	fmt.Println(myToDoItem)

}

func testFromJsonString() {
	//Lets populate our slice, its a global so if its already populated
	//by another test we can ignore
	if len(ToDoList2) == 0 {
		populateToDoList2()
	}

	//Lets convert our slice to JSON
	//Note that the second parameter is an error, so we need to check
	//it
	jsonBytes, err := json.MarshalIndent(ToDoList2, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	jsonString := string(jsonBytes)

	//Lets print the JSON
	fmt.Println("Lets see that we have a starting string")
	fmt.Println(jsonString)

	//Now lets convert it back to a slice of structs
	var myToDoList []ToDoItem2
	err = json.Unmarshal([]byte(jsonString), &myToDoList)
	if err != nil {
		fmt.Println(err)
	}
	//Lets print the slice of structs and make sure all is OK
	for idx, item := range myToDoList {
		fmt.Println(idx, ":", item)
	}

}

func RunJsonDemo() {
	fmt.Println("------ Running JSON Demo ------")
	testToDoToJson()
	testToDoToJsonPrettyPrint()
	testAnnotatedJson()
	testFromJsonString()
	fmt.Printf("-----------------------------------\n\n")
}
