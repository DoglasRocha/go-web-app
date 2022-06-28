package main

import (
	"html/template"
	"loja/models"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := produtos.BuscaTodosOsProdutos()

	templates.ExecuteTemplate(w, "Index", todosOsProdutos)
}
