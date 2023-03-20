package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// parsedTemplate := template.ParseFiles("./templates/" + tmpl)
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {

		fmt.Println("Error parsing template:", err)
		return
	}
}
