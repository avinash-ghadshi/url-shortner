package common

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type AppConfig struct {
	App struct {
		Server string   `json:"server"`
		Port   string   `json:"port"`
		Proxy  []string `json:"proxy"`
		Expiry int64    `json:"defaultexpiry"`
	} `json:"app"`
}

var Config *AppConfig

func ReadConfig(f *os.File, conf *AppConfig) {

	// Read the JSON file
	byteValue, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(byteValue, conf)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	Config = conf
}
