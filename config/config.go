package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
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
	fmt.Println("Reading config.json...")

	// 1. Try loading from executable path
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalln("Failed to get executable path:", err)
	}
	exeDir := filepath.Dir(exePath)
	configPath := filepath.Join(exeDir, "config.json")

	// Try to open from executable directory
	jsonFile, err := os.Open(configPath)
	if err != nil {
		// 2. Fallback: Try current working directory
		fmt.Println("Could not open config.json from executable directory, trying current directory...")

		configPath = "config.json"
		jsonFile, err = os.Open(configPath)
		if err != nil {
			log.Fatalf("Failed to open config.json from both executable and working directory: %v", err)
		}
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln("Error reading config.json:", err)
	}

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatalln("Error parsing config.json:", err)
	}

	fmt.Println("Config loaded from:", configPath)
}



// Getter function to access the config in main.go
func GetConfig() Config {
	return config
}
