package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Server     string
	PrivateKey string
}

func NewConfig(filename string) (Configuration, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Configuration{}, err
	}

	decoder := json.NewDecoder(file)
	var config Configuration
	err = decoder.Decode(&config)
	if err != nil {
		return Configuration{}, err
	}

	return config, nil
}

func main() {
	config, err := NewConfig("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(config)
}
