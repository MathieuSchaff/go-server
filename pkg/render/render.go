package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// parsedTemplate := template.ParseFiles("./templates/" + tmpl)
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {

		fmt.Println("Error parsing template:", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplateTest(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache
	_, inMap := tc[t]
	if !inMap {
		// need to create the template
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("using template from cache")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// add the template to the cache
	tc[t] = tmpl
	return nil
}
