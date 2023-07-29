# MongoDB CRUD

This is an example project that shows the basics of working with the MongoDB database in golang. Go provides an abstract datalayer using the go.mongodb.org/mongo-driver/ standard library.

## Prerequisites

Local MongoDB instance
Valid Golang installation

# Getting started

First you need to install the MongoDB Go drivers.

go get go.mongodb.org/mongo-driver
Or install it using a the dep package manager.

dep ensure --add go.mongodb.org/mongo-driver/mongo \
go.mongodb.org/mongo-driver/bson \
go.mongodb.org/mongo-driver/mongo/options
You can then start the application using go run.

go run mongodb.go
## Implementation
MongoDB CRUD
