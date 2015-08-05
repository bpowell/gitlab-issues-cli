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

type Project struct {
	Creator_id int
	Id         int
	Name       string
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

func GetAllProjects(url string) ([]Project, error) {
	var projects []Project
	r, err := http.Get(url)
	body, err := ioutil.ReadAll(r.Body)

	var list []map[string]interface{}
	err = json.Unmarshal(body, &list)
	if err != nil {
		return projects, err
	}

	for _, data := range list {
		d, _ := json.Marshal(data)
		var p Project
		err = json.Unmarshal(d, &p)
		if err != nil {
			return projects, err
		}

		projects = append(projects, p)
	}

	return projects, nil
}

func main() {
	config, err := NewConfig("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(config)

	url := fmt.Sprintf("%s%s?%s%s", config.Server, APIPATH, PRIVKEY, config.PrivateKey)
	projects, err := GetAllProjects(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(projects)
}
