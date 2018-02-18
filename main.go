package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/lucaswhitman/library-api/app"
)

type Configuration struct {
	Application struct {
		Port int `json:"port"`
	} `json:"application"`
	Database struct {
		Host         string `json:"host"`
		Port         int    `json:"port"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		DatabaseName string `json:"databaseName"`
	} `json:"database"`
}

func main() {
	a := app.App{}

	conf, err := getConf("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	a.Initialize(conf.Database.Host, conf.Database.Port, conf.Database.Username, conf.Database.Password, conf.Database.DatabaseName)

	a.Run(conf.Application.Port)
}

func getConf(fileName string) (Configuration, error) {
	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)
	conf := Configuration{}
	err := decoder.Decode(&conf)
	return conf, err
}
