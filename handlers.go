package main

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func Hangman(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "hangman")
}

func Settings(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "settings")
}

func renderTemplate(w http.ResponseWriter, html string) {
	t, err := template.ParseFiles("./templates/html/" + html + ".html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
