package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"drexel.edu/todo/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Global variables to hold the command line flags to drive the todo CLI
// application
var (
	hostFlag string
	portFlag uint
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
func processCmdLineFlags() {

	//Note some networking lingo, some frameworks start the server on localhost
	//this is a local-only interface and is fine for testing but its not accessible
	//from other machines.  To make the server accessible from other machines, we
	//need to listen on an interface, that could be an IP address, but modern
	//cloud servers may have multiple network interfaces for scale.  With TCP/IP
	//the address 0.0.0.0 instructs the network stack to listen on all interfaces
	//We set this up as a flag so that we can overwrite it on the command line if
	//needed
	flag.StringVar(&hostFlag, "h", "0.0.0.0", "Listen on all interfaces")
	flag.UintVar(&portFlag, "p", 1080, "Default Port")

	flag.Parse()
}

// main is the entry point for our todo API application.  It processes
// the command line flags and then uses the db package to perform the
// requested operation
func main() {
	processCmdLineFlags()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())

	apiHandler, err := api.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//HTTP Standards for "REST" APIS
	//GET - Read/Query
	//POST - Create
	//PUT - Update
	//DELETE - Delete

	app.Get("/todo", apiHandler.ListAllTodos)
	app.Post("/todo", apiHandler.AddToDo)
	app.Put("/todo", apiHandler.UpdateToDo)
	app.Delete("/todo", apiHandler.DeleteAllToDo)
	app.Delete("/todo/:id<int>", apiHandler.DeleteToDo)
	app.Get("/todo/:id<int>", apiHandler.GetToDo)

	app.Get("/crash", apiHandler.CrashSim)
	app.Get("/crash2", apiHandler.CrashSim2)
	app.Get("/crash3", apiHandler.CrashSim3)
	app.Get("/health", apiHandler.HealthCheck)

	//We will now show a common way to version an API and add a new
	//version of an API handler under /v2.  This new API will support
	//a path parameter to search for todos based on a status
	v2 := app.Group("/v2")
	v2.Get("/todo", apiHandler.ListSelectTodos)

	serverPath := fmt.Sprintf("%s:%d", hostFlag, portFlag)
	log.Println("Starting server on ", serverPath)
	app.Listen(serverPath)
}
