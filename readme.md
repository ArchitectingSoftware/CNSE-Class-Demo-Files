## Cloud Native Software Engineering

This repository will be used for in class demonstrations and to provide students with scaffolded code for my cloud native software engineering class.

Contents:

1. [Go Tutorial](./gotutorial/).  This directory contains a basic go language tutorial to get you ready for this class
2. [ToDo Application](./todo/).  This directory contains the assignment for building a simple `todo` app.  It has a significant amount of scaffolded code, and is a good initial example in building CLI-based applications in Go.
3. [ToDo API (Demo)](./todo-api/).  This directory contains a demo/technical tutorial that we will be using to explore creating APIs in go
4. [GCP IaaS (Demo)](./infrastructure-automation/).  This directory contains a demo/technical tutorial on using automation to create a virtual machine in the cloud and push some code to it. There are 2 sub-demos, one showing the use of Terraform, which is an industry leading automation tool, and the other using Pulumi, that embraces using traditional programming languages, versus a custom configuration-as-code format.
5. [ToDo API With Events (Demo)](./todo-api-w-events/).  This directory an extension of the basic `todo-api`.  It illustrates `goroutines`, `channels`, and `events`