SHELL := /bin/bash

.PHONY: help
help:
	@echo "Usage make <TARGET>"
	@echo ""
	@echo "  Targets:"
	@echo "	   build				Build the todo executable"
	@echo "	   run					Run the todo program from code"
	@echo "	   run-bin				Run the todo executable"
	@echo "	   test					Run the tests"
	@echo "	   test-verbose			Run the tests with verbose output"
	@echo "	   restore-db			Restore the sample database (unix/mac)"
	@echo "	   restore-db-windows	Restore the sample database (windows)"
	@echo "	   add-sample			Add a sample row"


.PHONY: build
build:
	go build .

.PHONY: run
run:
	go run main.go

.PHONY: run-bin
run-bin:
	./todo

.PHONY: restore-db
restore-db:
	(cp ./data/todo.json.bak ./data/todo.json)

.PHONY: restore-db-windows
restore-db-windows:
	(copy.\data\todo.json.bak .\data\todo.json)

.PHONY: test
test:
	go test ./tests

.PHONY: test-verbose
test-verbose:
	go test ./tests -v

.PHONY: add-sample
add-sample:
	go run main.go -a '{ "id":99, "title":"sample item", "done":true}'
