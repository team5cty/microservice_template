package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

/*Yaml has 2 fields currently
-Database
-Endpoints array - list of endpoints
*/

type Yaml struct {
	Database struct {
		Provider string `yaml:"provider"`
		Url      string `yaml:"url"`
	} `yaml:"database"`

	Endpoints []Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	path   string `yaml:"string"`
	method string `yaml:"method"`

	Schema struct {
		Type       string            `yaml:"type"`
		Properties map[string]string `yaml:"properties"`
	}
}

func main() {
	yamlfile, err := os.ReadFile("example.yaml")
	if err != nil {
		fmt.Printf("Failed to read input.yaml file: %s\n", err.Error())
		return
	}
	yamlobject := &Yaml{}
	err = yaml.Unmarshal(yamlfile, yamlobject)
	if err != nil {
		fmt.Printf("Failed to parse YAML file: %s\n", err.Error())
	}

	fmt.Println(yamlobject)

}
