package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)



// Config struct to match the JSON structure
type Config struct {
	DbServer struct {
		Driver           string `json:"driver"`
		ConnectionString string `json:"connectionstring"`
	} `json:"dbserver"`
	SapServer struct {
		Username        string `json:"username"`
		Password        string `json:"password"`
		Companydb       string `json:"companydb"`
		Servicelayerurl string `json:"servicelayerurl"`
	} `json:"sapserver"`
	Port          int `json:"port"`
	WebSocketPort int `json:"websocketPort"`
}

// Function to read the config file
var config Config

func ReadConfig() {
	// Open our jsonFile
	fmt.Println("Fetching config data")
	jsonFile, err := os.Open("config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println("error opening config.json", err)
		log.Fatalln("error opening config.json", err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	bytevalue, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(bytevalue, &config)
	if err != nil {
		fmt.Println("Error reading config file", err)
		log.Fatalln("Error reading config file", err)
	} else {
		fmt.Println("Config data", config)
	}

}

// Getter function to access the config in main.go
func GetConfig() Config {
	return config
}
