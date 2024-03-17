package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"

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
	Path   string `yaml:"path"`
	Method string `yaml:"method"`

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
	tmpl := `
	package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

{{range .Endpoints}}
func {{.Method | toLower}}{{.Path | clean}}Handler(w http.ResponseWriter, r *http.Request) {
	// Implement logic for {{.Method}} {{.Path}}
	fmt.Fprintf(w, "Handling %s request for %s", r.Method, r.URL.Path)
}

{{end}}
func main() {
	r := mux.NewRouter()

	{{range .Endpoints}}
	r.HandleFunc("{{.Path}}", {{.Method | toLower}}{{.Path | clean}}Handler).Methods("{{.Method}}")
	{{end}}

	fmt.Println("Server is running...")
	http.ListenAndServe(":8080", r)
}
`

	// Create a new template and parse the template string
	t := template.Must(template.New("restAPI").Funcs(template.FuncMap{"clean": cleanMethodName, "toLower": strings.ToLower}).Parse(tmpl))

	// Execute the template with the provided data
	err = t.Execute(os.Stdout, yamlobject)
	if err != nil {
		fmt.Printf("Error executing template: %s\n", err.Error())
	}
}

func cleanMethodName(s string) string {
	s = strings.ReplaceAll(s, "/", "")
	s = strings.ReplaceAll(s, "{", "")
	s = strings.ReplaceAll(s, "}", "")
	return strings.ToLower(s)
}
