# music-service
Implementation of a REST API communicating with a relational database. Uses Gorm library for DB access and test containers library for integration tests

## Build and Run:
`go mod vendor`

`docker build -t neiln3121/music-service .`

Run with MS SQL Server using

`docker-compose up`

Once mssql server is running, you can load data using the loader package:

`go run loader/cmd/main.go`

Once the application is running, all artists can be accessed from:

http://localhost:8080/api/artists

An album can be retreived using:

http://localhost:8080/api/albums/1

There are 3 artists, 6 albums and a total of 63 tracks inserted by the loader package. This package is also called during the integration tests. These spin up a mssql server using the testcontainers library.

TODO: Post requests and authorisation middleware