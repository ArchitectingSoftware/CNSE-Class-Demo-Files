#!/bin/bash
#IMPORT CONTAINERS THAT ARE LOCAL - YOU CAN AVOID THIS IF YOU PUSH YOUR CONTAINERS TO A REGISTRY LIKE DOCKERHUB
kind load docker-image --name cnse-class architectingsoftware/cnse-pub-api:v1
kind load docker-image --name cnse-class architectingsoftware/cnse-publist-api:v1