## Basic Go Tutorial

This project demonstrates some basic aspects of the go language.

To start, in an empty directory crate a go module `go mod init architectingsoftware.com/gotutorial`.  Of course, if you clone this project from GitHub you will not have to do this.

The `go.mod` file contains the initial contents when created:

```
module architectingsoftware.com/gotutorial

go 1.20
```

This names the module and specifies the version of the go compiler to be used to compile the program.  As you add dependencies to your program, the `go.mod` file will be updated automatically.  Once you start adding dependencies, the go compiler will also create the `go.sum` file that not only has your dependencies, but the transitive dependencies of your dependencies as well.  You dont need to worry about these files, they are for dependency management. That said, its good to know what these things are, as the `go.sum` file not only has all of your program dependencies, its keeps track of versions and includes sha checksums to make sure that you have the correct versions that have not been tampered with.

