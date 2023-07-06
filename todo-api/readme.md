## ToDo API Demo

This is a demo application showing many aspects of how to use the Golang Gin
framework to create an API.

It keeps `todo` items in memory for this demo.  The makefile allows you to 
exercise the API.  For example you can load the database, query by item,
and so on.

To see everything you can do you can just run `make` and get help.  See below.  Also notice that some of the make targets take parameters.  To do this you add a key=value on the `make` command line.  For example, to get a `todo` with an id of `2`. you run `make id=2 get-by-id`

```
➜  todo-api git:(main) make
Usage make <TARGET>

  Targets:
           build                        Build the todo executable
           run                          Run the todo program from code
           run-bin                      Run the todo executable
           load-db                      Add sample data via curl
           get-by-id                    Get a todo by id pass id=<id> on command line
           get-all                      Get all todos
           update-2                     Update record 2, pass a new title in using title=<title> on command line
           delete-all                   Delete all todos
           delete-by-id                 Delete a todo by id pass id=<id> on command line
           get-v2                       Get all todos by done status pass done=<true|false> on command line
           get-v2-all                   Get all todos using version 2
```

### Why use the gin framework?

Many people in the golang community are opposed to using frameworks because the standard library provides robust function out-of-the-box.  However, the golang gin framework reduces a lot of the code you need to write and has a lot of nice features out of the box.  As far as I know its still the most popular and widely used API framework for go.

Online documentation for gin can be found here:

1. GitHub page: https://github.com/gin-gonic/gin
2. Go Docs: https://pkg.go.dev/github.com/gin-gonic/gin?utm_source=godoc
3. Gin homepage: https://gin-gonic.com/