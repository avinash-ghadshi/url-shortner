package main

import (
	"fmt"
	"os"
	"urlshortener/server/libs/common"
)

var Config common.AppConfig

func init() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Printf("Error loading config file: %v", err.Error())
		os.Exit(1)
	}
	common.ReadConfig(jsonFile, &Config)
	jsonFile.Close()
}

func main() {
	app := setupRouter()
	app.Run(Config.App.Server + ":" + Config.App.Port)
}
