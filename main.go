package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/lucaswhitman/library-api/app"
)

type Configuration struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"databaseName"`
}

func main() {
	a := app.App{}

	conf := getConf("./config.json")

	a.Initialize(conf.Host, conf.Port, conf.Username, conf.Password, conf.DatabaseName)

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
