package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		input := r.URL.Query().Get("input")
		renderTemplate(w, "<h1>"+input+"</h1>")
	})

	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, content string) {
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Sample Page</title>
	</head>
	<body>
		{{.}}
	</body>
	</html>
	`
	t := template.Must(template.New("index").Parse(tmpl))
	t.Execute(w, template.HTML(content))
}

