# microservice_template
w.Header().Set("Content-Type", "application/json")
var {{.endpoint.Name | tolower}} {{.endpoint.Name}}
// Implement logic for {{.endpoint.Method}} {{.endpoint.Path}}
{{.endpoint.Name | tolower}}.ToJSON(w)