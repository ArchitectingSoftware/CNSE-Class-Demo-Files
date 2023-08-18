#!/bin/bash
curl -d '{ "id": 1, "title": "Learn Go / GoLang", "done": false }' -H "Content-Type: application/json" -X POST http://localhost:1080/todo 
curl -d '{ "id": 2, "title": "Learn Kubernetes", "done": false}' -H "Content-Type: application/json" -X POST http://localhost:1080/todo 
curl -d '{"id": 3,"title": "Learn Cloud Native Architecure","done":false}' -H "Content-Type: application/json" -X POST http://localhost:1080/todo
curl -d '{"id": 100,"title": "Office hours are helpful","done":false}' -H "Content-Type: application/json" -X POST http://localhost:1080/todo