#notes
* tests could probably be refactored a bit to reduce the bulk of the copy-pasta code

#dependencies
go get github.com/stretchr/testify/assert
go get github.com/gorilla/mux github.com/lib/pq

#requirements
Go
Postgresql

#instructions
run database migration

# Tests
To run tests go test ./...