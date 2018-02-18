package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/lucaswhitman/library-api/app"
)

type Configuration struct {
	host         string
	port         int
	username     string
	password     string
	databaseName string
}

func main() {
	a := app.App{}

	conf := getConf("./config.json")

	a.Initialize(conf.host, conf.port, conf.username, conf.password, conf.databaseName)

	a.Run(":8080")
}

func getConf(fileName string) Configuration {
	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)
	conf := Configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		log.Fatal(err)
	}
	return conf
}
