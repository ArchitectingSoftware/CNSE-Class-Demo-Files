## ToDo App

This is the class assignment for the ToDo list CLI assignment.
**READ THIS CAREFULLY AS THESE ARE THE DIRECTIONS**

Most cloud native applications are packaged as server side code, and operated on with command line interface (CLI) tools.   For our first programming assignment we are going to implement a go language CLI tool to manage a list of todo items.

This application will be driven by a simple text file based database.  See the `todo.json` file.  Notice that this file is structured as a JSON array, with collections of individual JSON objects.  Each object contains an `id`, `description` and a `done` flag.  For example:

```
[
  {
    "id": 1,
    "title": "Learn Go / GoLang",
    "done": false
  },
  {
    "id": 2,
    "title": "Learn Kubernetes",
    "done": false
  }
]
```
By default our program uses `./data/todo.json` as the default database.  You can override the database name from the command line via the `-db` flag providing a new database name.  For example `-db ./data/my_new_database.db`.  More on that later. 

### What you need to do

Carefully study the provided code.  Its a helpful scaffold. The code should run as is, albeit it does not do very much.  Within the code you will see a number of comments that look like:

```
// TODO: <What you need to do>
```

Some of the `TODO:` prompts involve you writing comments to answer specific questions, others are prompts describing the code you are expected to develop.

Answer all of the `TODO:` items and then upload your code to your GitHub/GitLab repository.  On blackboard, send me a link to your solution. 

Note that there are also some `TODO:` items marked as extra credit. **You do not have to do any of these items if you do not want any extra credit.  That said, I provided them less for you to get extra credit, and more for helping you to grasp a deeper understanding of go once you solve the basic assignment requirements**

Remember from our first lecture that you will only have one repo this entire term for many different deliverables.  Please place your solution under the `/todo` directory in your repo

To make your life easier I also am providing a makefile to automate a lot of the common commands your will be using.  You can thank me later.

In most of the other assignments I will also be requiring you to create a readme file in markdown and will ask for specific information about how to
use your code.

There is no need to do that with this assignment, as the command line options are fixed based on the scaffold that I provided.  Thus the command line options for this CLI should be (and should not be changed by you):

```
todo git:(main) ✗ ./todo -h
Usage of ./todo:
  -a string
        Add an item to the database
  -d int
        Delete an item from the database
  -db string
        Name of the database file (default "./data/todo.json")
  -l    List all the items in the database
  -q int
        Query an item in the database
  -s    Change item 'done' status to true or false
  -u string
        Update an item in the database
  ```


