package main

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v3"
)

/*Yaml has 3 fields
-Module - name of module to be generated
-Database - url and schema of tables
-Endpoints array - list of endpoints
*/

type Yaml struct {
	Module string `yaml:"module"`

	Database struct {
		Provider string  `yaml:"provider"`
		Url      string  `yaml:"url"`
		Models   []Model `yaml:"models"`
	} `yaml:"database"`

	Endpoints []Endpoint `yaml:"endpoints"`
}

type Model struct {
	Table  string            `yaml:"table"`
	Schema map[string]string `yaml:"schema"`
}

type Endpoint struct {
	Name   string `yaml:"name"`
	Path   string `yaml:"path"`
	Method string `yaml:"method"`
	Table  string `yaml:"table"`
	Json   struct {
		Type       string            `yaml:"type"`
		Properties map[string]string `yaml:"properties"`
	} `yaml:"json"`
}

func main() {

	//read example.yaml and parse it into yamlobject
	yamlfile, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to read input.yaml file: %s\n", err.Error())
		return
	}
	yamlobject := &Yaml{}
	err = yaml.Unmarshal(yamlfile, yamlobject)
	if err != nil {
		fmt.Printf("Failed to parse YAML file: %s\n", err.Error())
	}

	//Create a new git module with name as per in YAML
	//Inside it create handlers folder which will contain handler function and struct of one endpoint
	//Create prisma folder inside which there will be schema.prisma file.
	//Run command go mod init and go get to install packages packages gorilla/mux and lib/pq and prisma.

	os.Mkdir(yamlobject.Module, os.ModePerm)
	os.Chdir(yamlobject.Module)
	os.Mkdir("handlers", os.ModePerm)
	os.Mkdir("prisma", os.ModePerm)
	cmd := exec.Command("go", "mod", "init", yamlobject.Module)
	cmd.Run()
	cmd = exec.Command("go", "get", "github.com/gorilla/mux")
	cmd.Run()
	cmd = exec.Command("go", "get", "github.com/steebchen/prisma-client-go")
	cmd.Run()
	os.Chdir("..")

	//There are three template files:-
	//main - for generating main.go
	//handler - for generating files inside handler folder
	//database - for generating file inside database folder

	//read main template and generate main.go inside module
	template_file_buffer, err := os.ReadFile("templates/main")
	if err != nil {
		fmt.Printf("Failed to read template.go file: %s\n", err.Error())
	}
	template_output_buffer, err := os.Create(yamlobject.Module + "/main.go")
	if err != nil {
		fmt.Printf("Failed to create output.go file: %s\n", err.Error())
	}
	t := template.Must(template.New("restAPI").Funcs(template.FuncMap{"tolower": strings.ToLower}).Parse(string(template_file_buffer)))
	err = t.Execute(template_output_buffer, yamlobject)
	if err != nil {
		fmt.Printf("Error executing template: %s\n", err.Error())
	}
	template_file_buffer, err = os.ReadFile("templates/prisma")
	if err != nil {
		fmt.Printf("Failed to read template.go file: %s\n", err.Error())
	}
	template_output_buffer, err = os.Create(yamlobject.Module + "/prisma/schema.prisma")
	if err != nil {
		fmt.Printf("Failed to create output.go file: %s\n", err.Error())
	}
	t = template.Must(template.New("restAPI").Parse(string(template_file_buffer)))
	err = t.Execute(template_output_buffer, yamlobject)
	if err != nil {
		fmt.Printf("Error executing template: %s\n", err.Error())
	}
	os.Chdir(yamlobject.Module + "/prisma")
	fmt.Println(os.Getwd())
	cmd = exec.Command("go", "run", "github.com/steebchen/prisma-client-go", "db", "push")

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error occurred:", err)
		return
	}

	os.Chdir("..")
	os.Chdir("..")

	//Loop through each endpoints and for each, generate a file
	// inside handlers folder in module using handler template
	for _, i := range yamlobject.Endpoints {
		template_file_buffer, err := os.ReadFile("templates/handlers")
		if err != nil {
			fmt.Printf("Failed to read template.go file: %s\n", err.Error())
		}
		template_output_buffer, err := os.Create(yamlobject.Module + "/handlers/" + i.Name + ".go")
		if err != nil {
			fmt.Printf("Failed to create output.go file: %s\n", err.Error())
		}
		t := template.Must(template.New("restAPI").Funcs(template.FuncMap{
			"tolower":      strings.ToLower,
			"isGET":        isGET,
			"isPOST":       isPOST,
			"title":        strings.Title,
			"hasPathParam": hasPathParam,
			"isList":       isList,
		}).Parse(string(template_file_buffer)))
		//We need name of module as variable inside handler template,
		// so passing this map with both a endpoint and module name.
		data := map[string]any{
			"endpoint": i,
			"module":   yamlobject.Module,
		}

		err = t.Execute(template_output_buffer, data)
		if err != nil {
			fmt.Printf("Error executing template: %s\n", err.Error())
		}
	}
}

// functions for checking if method is get or post
// inorder to add either ToJSON or FromJSON methods to struct in handlers template.
func isGET(s string) bool {
	if s == "GET" {
		return true
	}
	return false
}

func isPOST(s string) bool {
	if s == "POST" {
		return true
	}
	return false
}
func hasPathParam(path string) bool {
	return strings.Contains(path, "{")
}

func isList(st string) bool {
	if st == "list" {
		return true
	}
	return false
}
