package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Configuration struct {
	Server     string
	PrivateKey string
}

const (
	APIPATH = "/api/v3/projects"
	PRIVKEY = "private_token="
)

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

	r, err := http.Get(fmt.Sprintf("%s%s?%s%s", config.Server, APIPATH, PRIVKEY, config.PrivateKey))
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(body)
}
