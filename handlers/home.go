package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

const homePath = "/"

// homeHandler is the home page handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/home.page.tmpl")
	if err != nil {
		fmt.Printf("error parsing template: %s", err.Error())
	}

	err = template.Execute(w, nil)
	if err != nil {
		fmt.Printf("error executing template: %s", err.Error())
		return
	}
}

func Home() {
	http.HandleFunc(homePath, homeHandler)
}