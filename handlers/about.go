package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

const aboutPath = "/about"

// aboutHandler is the about page handler
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/about.page.tmpl")
	if err != nil {
		fmt.Printf("error parsing template: %s", err.Error())
	}

	err = template.Execute(w, nil)
	if err != nil {
		fmt.Printf("error executing template: %s", err.Error())
		return
	}
}

func About() {
	http.HandleFunc(aboutPath, aboutHandler)
}