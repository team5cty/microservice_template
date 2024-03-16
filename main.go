package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

/*Yaml has 3 fields currently
-Database
-Model- not implemented
-Endpoint
*/

type Yaml struct {
	Database struct {
		Provider string `yaml:"provider"`
		Url      string `yaml:"url"`
	} `yaml:"database"`

	//Models []Model `yaml:"models"`

	Endpoints []Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	path   string `yaml:"string"`
	method string `yaml:"method"`

	Schema struct {
		//Model Model `yaml:"model"`
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
	yaml.Unmarshal(yamlfile, yamlobject)
	fmt.Println(yamlobject)
}
