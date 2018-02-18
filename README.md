# Library-API
This is a sample API built with the Go programming language. It is built on top of PostgreSQL and GMux

## Requirements
In order to run this application locally you'll need to have a PostgreSQL instance, and Go installed. 

## Local Development
### Configuration
Some base configuration is configured, this can be changed if your local environment requires it
```javascript
{
  "application": {
    "port": 8000
  },
  "database": {
    "host": "127.0.0.1",
    "port": 5432,
    "username": "library_test",
    "password": "password",
    "databaseName": "library_test"
  }
}
```
There is both a `config.json` for the main app and a `config_test.json` for running tests to keep the data separate. *The tests will clear all data from the tables so don't use the same database!*

### Dependencies
The following dependencies are needed and can be installed with the commands:
```bash
go get github.com/stretchr/testify/assert
go get github.com/gorilla/mux github.com/lib/pq
```
There is also an initial database migration script for spinning up the databases and roles needed that can be ran with the following command:
```bash
./initial_database_migration.sh
```
In the future this should probably be moved into a database migration tool...

### Building
It was developed on a Mac, but Go should be able to compile the proper binaries using the following command:
```bash
go build
```

### Running Tests
To run tests:
```bash
 go test ./...
```
