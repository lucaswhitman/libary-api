# Library-API
This is a sample API built with the Go programming language. It is built on top of PostgreSQL and GMux. It is configured for Docker and Kubernetes.

## Requirements
In order to run this application locally you'll need to have a PostgreSQL instance, and Go installed. 

## Local Development
### Configuration
Some base configuration is configured, this can be changed if your local environment requires it. *Base file is configured to run in Docker.*
```javascript
{
  "application": {
    "port": 80
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
make build
```

### Running Tests
To run tests:
```bash
make test
```

## Using Docker locally
Requires that Docker is installed on the local machine
```bash
make local
```

## Testing deployment with Minikube
Setting up Minikube (assumes installed)
```bash
eval $(minikube docker-env)
minikube start
```
Build binary for Linux
```bash
make build-linux
```
Create kubernetes cluster from .yml file and expose via a load balancer
```bash
kubectl create -f kompose.yml
kubectl expose deployment app --type=LoadBalancer
```
The --type=LoadBalancer flag indicates that you want to expose your Service outside of the cluster. On cloud providers that support load balancers, an external IP address would be provisioned to access the Service. On Minikube, the LoadBalancer type makes the Service accessible through the minikube service command.
```bash
minikube service app
```
This will open a browser, you can then use this address and port to access the api:
```
<ip:port>/books
```

cleanup (probably don't want to do this in production...):
```bash
kubectl delete pods --all && kubectl delete service --all && kubectl delete deployments --all && kubectl delete persistentvolumeclaims --all
```

## Sample usage
This app can be tested manually via PostMan or cURL, here are some information to get you started:

### The Book Model
All fields are required, new books should not have an id set.
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
`POST /books`
#### Example
```bash
curl -i -H "Content-Type: application/json" -X POST -d '{
  "title":"Hitchhikers Guide to the Galaxy",
  "author": "Douglas Adams",
  "publisher": "Pan Books",
  "publishDate": "1979-10-12T11:45:26.371Z",
  "rating": 3,
  "status": "CheckedOut"
}' http://localhost:80/books
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
`GET /books/{id}`
#### Example:
```bash
curl -i -H "Content-Type: application/json" -X GET http://localhost:80/books/1
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
`GET /books`
#### Example:
```bash
curl -i -H "Content-Type: application/json" -X GET http://localhost:80/books
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
`DELETE /books/{id}`
#### Example:
```bash
curl -i -H "Content-Type: application/json" -X DELETE http://localhost:80/books/1
```
#### Returns:
The status of the request
```javascript
{
  "result":"success"
}
```

## Future design thoughts
* Increased metadata - Things like author, review, and publisher should be first class citzens in the data model. This would allow for advanced querying such as getting all high rated books by your favorite author.
* Authentication - Certain APIs should be restricted to logged in users. This would allow for restricting the ability to delete books, and linking comments to users.
* Increased testing - All the major endpoints are covered, but if we were taking this to production I'd want to add additional unit/integration tests. I'd also likely want to add some additonal automated testing via an external test service.
* UI - If we want to expose this service to end users we'd likely need a UI.
* Free text search and facets - PostgreSQL has the ability to allow for free text search, we should leverage it in the API to allow the user to search for terms in the title or author. We could also expose metadata as facets to allow the user to interactively browse through the data.
* More robust deployments - The process for deploying this app is pretty manual at this point. If I had to do it more than once I'd likely want to set up a CI/CD pipeline via a service like Jenkins or CircleCI.
