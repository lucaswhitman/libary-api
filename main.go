package main

import (
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	host         string
	port         int
	username     string
	password     string
	databaseName string
}

func main() {
	a := App{}

	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		log.Fatal(err)
	}

	a.Initialize(conf.host, conf.port, conf.username, conf.password, conf.databaseName)

	a.Run(":8080")
}
