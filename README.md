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

## Sample usage
This app can be tested manually via PostMan or cURL, here are some information to get you started:

### The Book Model
```javascript
{
  "id": 1,
  "title":"Hitchhikers Guide to the Galaxy",
  "author": "Douglas Adams",
  "publisher": "Pan Books",
  "publishDate": "1979-10-12T11:45:26.371Z",
  "rating": 3,
  "status": "CheckedOut"
}
```

### Create a New Note
POST /books
BODY a book
#### Example
```bash
curl -i -H "Content-Type: application/json" -X POST -d '{
  "title":"Hitchhikers Guide to the Galaxy",
  "author": "Douglas Adams",
  "publisher": "Pan Books",
  "publishDate": "1979-10-12T11:45:26.371Z",
  "rating": 3,
  "status": "CheckedOut"
}' http://localhost:8000/books
```
#### Returns:
A saved book...
```javascript
{
  "id":1,
  "title":"Hitchhikers Guide to the Galaxy",
  "author":"Douglas Adams",
  "publisher":"Pan Books",
  "publishDate":"1979-10-12T11:45:26.371Z",
  "rating":3,
  "status":"CheckedOut"
}
```

### Get an existing book
You can get a book using an API call:
GET /books/{id}
#### Example:
```bash
curl -i -H "Content-Type: application/json" -X GET http://localhost:8000/books/1
```
#### Returns:
The requested book
```javascript
{
  "id":1,
  "title":"Hitchhikers Guide to the Galaxy",
  "author":"Douglas Adams",
  "publisher":"Pan Books",
  "publishDate":"1979-10-12T11:45:26.371Z",
  "rating":3,
  "status":"CheckedOut"
}
```

### Get all books
I can get all books using an API call:
GET /books
#### Example:
```bash
curl -i -H "Content-Type: application/json" -X GET http://localhost:8000/books
```
Returns:
A list of books
```javascript
[
  {
    "id":1,
    "title":"Hitchhikers Guide to the Galaxy",
    "author":"Douglas Adams",
    "publisher":"Pan Books",
    "publishDate":"1979-10-12T00:00:00Z",
    "rating":3,"status":"CheckedOut"
  },{
    "id":2,
    "title":"The Great Book of Amber",
    "author":"Roger Zelazny",
    "publisher":"Harper Voyager",
    "publishDate":"1979-10-12T00:00:00Z",
    "rating":3,"status":"CheckedOut"
  }
]
```

### Delete an existing book
You can delete a book using an API call:
DELETE /books/{id}
#### Example:
```bash
curl -i -H "Content-Type: application/json" -X DELETE http://localhost:8000/books/1
```
#### Returns:
The status of the request
```javascript
{
  "result":"success"
}
```




