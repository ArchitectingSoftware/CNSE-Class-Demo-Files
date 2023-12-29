package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"drexel.edu/todo/db"
)

// Global variables to hold the command line flags to drive the todo CLI
// application
var (
	dbFileNameFlag string
	restoreDbFlag  bool
	listFlag       bool
	itemStatusFlag bool
	queryFlag      int
	addFlag        string
	updateFlag     string
	deleteFlag     int
)

type AppOptType int

// To make the code a little more clean we will use the following
// constants as basically enumerations for different options.  This
// allows us to use a switch statement in main to process the command line
// flags effectively
const (
	LIST_DB_ITEM AppOptType = iota
	RESTORE_DB_ITEM
	QUERY_DB_ITEM
	ADD_DB_ITEM
	UPDATE_DB_ITEM
	DELETE_DB_ITEM
	CHANGE_ITEM_STATUS
	NOT_IMPLEMENTED
	INVALID_APP_OPT
)

// processCmdLineFlags parses the command line flags for our CLI
//
// TODO: This function uses the flag package to parse the command line
//		 flags.  The flag package is not very flexible and can lead to
//		 some confusing code.

//			 REQUIRED:     Study the code below, and make sure you understand
//						   how it works.  Go online and readup on how the
//						   flag package works.  Then, write a nice comment
//				  		   block to document this function that highights that
//						   you understand how it works.
//
//			 EXTRA CREDIT: The best CLI and command line processor for
//						   go is called Cobra.  Refactor this function to
//						   use it.  See github.com/spf13/cobra for information
//						   on how to use it.
//
//	 YOUR ANSWER: <GOES HERE>
func processCmdLineFlags() (AppOptType, error) {
	flag.StringVar(&dbFileNameFlag, "db", "./data/todo.json", "Name of the database file")
	flag.BoolVar(&restoreDbFlag, "restore", false, "Restore the database from the backup file")
	flag.BoolVar(&listFlag, "l", false, "List all the items in the database")
	flag.IntVar(&queryFlag, "q", 0, "Query an item in the database")
	flag.StringVar(&addFlag, "a", "", "Add an item to the database")
	flag.StringVar(&updateFlag, "u", "", "Update an item in the database")
	flag.IntVar(&deleteFlag, "d", 0, "Delete an item from the database")
	flag.BoolVar(&itemStatusFlag, "s", false, "Change item 'done' status to true or false")

	flag.Parse()

	var appOpt AppOptType = INVALID_APP_OPT

	//show help if no flags are set
	if len(os.Args) == 1 {
		flag.Usage()
		return appOpt, errors.New("no flags were set")
	}

	// Loop over the flags and check which ones are set, set appOpt
	// accordingly
	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "l":
			appOpt = LIST_DB_ITEM
		case "restore":
			appOpt = RESTORE_DB_ITEM
		case "q":
			appOpt = QUERY_DB_ITEM
		case "a":
			appOpt = ADD_DB_ITEM
		case "u":
			appOpt = UPDATE_DB_ITEM
		case "d":
			appOpt = DELETE_DB_ITEM

		//TODO: EXTRA CREDIT - Implment the -s flag that changes the
		//done status of an item in the database.  For example -s=true
		//will set the done status for a particular item to true, and
		//-s=false will set the done states for a particular item to
		//false.
		//
		//HINT FOR EXTRA CREDIT
		//Note the -s option also requires an id for the item to that
		//you want to change.  I recommend you use the -q option to
		//specify the item id.  Therefore, the -s option is only valid
		//if the -q option is also set
		case "s":
			//For extra credit you will need to change some things here
			//and also in main under the CHANGE_ITEM_STATUS case
			appOpt = CHANGE_ITEM_STATUS
		default:
			appOpt = INVALID_APP_OPT
		}
	})

	if appOpt == INVALID_APP_OPT || appOpt == NOT_IMPLEMENTED {
		fmt.Println("Invalid option set or the desired option is not currently implemented")
		flag.Usage()
		return appOpt, errors.New("no flags or unimplemented were set")
	}

	return appOpt, nil
}

// main is the entry point for our todo CLI application.  It processes
// the command line flags and then uses the db package to perform the
// requested operation
func main() {

	//Process the command line flags
	opts, err := processCmdLineFlags()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Create a new db object
	todo, err := db.New(dbFileNameFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Switch over the command line flags and call the appropriate
	//function in the db package
	switch opts {
	case RESTORE_DB_ITEM:
		fmt.Println("Running RESTORE_DB_ITEM...")
		if err := todo.RestoreDB(); err != nil {
			fmt.Println("Error: ", err)
			break
		}
		fmt.Println("Database restored from backup file")
	case LIST_DB_ITEM:
		fmt.Println("Running QUERY_DB_ITEM...")
		todoList, err := todo.GetAllItems()
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
		for _, item := range todoList {
			todo.PrintItem(item)
		}
		fmt.Println("THERE ARE", len(todoList), "ITEMS IN THE DB")
		fmt.Println("Ok")

	case QUERY_DB_ITEM:
		fmt.Println("Running QUERY_DB_ITEM...")
		item, err := todo.GetItem(queryFlag)
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
		todo.PrintItem(item)
		fmt.Println("Ok")
	case ADD_DB_ITEM:
		fmt.Println("Running ADD_DB_ITEM...")
		item, err := todo.JsonToItem(addFlag)
		if err != nil {
			fmt.Println("Add option requires a valid JSON todo item string")
			fmt.Println("Error: ", err)
			break
		}
		if err := todo.AddItem(item); err != nil {
			fmt.Println("Error: ", err)
			break
		}
		fmt.Println("Ok")
	case UPDATE_DB_ITEM:
		fmt.Println("Running UPDATE_DB_ITEM...")
		item, err := todo.JsonToItem(updateFlag)
		if err != nil {
			fmt.Println("Update option requires a valid JSON todo item string")
			fmt.Println("Error: ", err)
			break
		}
		if err := todo.UpdateItem(item); err != nil {
			fmt.Println("Error: ", err)
			break
		}
		fmt.Println("Ok")
	case DELETE_DB_ITEM:
		fmt.Println("Running DELETE_DB_ITEM...")
		err := todo.DeleteItem(deleteFlag)
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
		fmt.Println("Ok")
	case CHANGE_ITEM_STATUS:
		//For the CHANGE_ITEM_STATUS extra credit you will also
		//need to add some code here
		fmt.Println("Running CHANGE_ITEM_STATUS...")
		fmt.Println("Not implemented yet, but it can be for extra credit")
		fmt.Println("Ok")
	default:
		fmt.Println("INVALID_APP_OPT")
	}
}
