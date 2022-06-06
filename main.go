package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Camiseta", "Azul, bem bonita", 39, 5},
		{"Tênis", "Confortável", 89, 3},
		{"Fone", "Muito bom", 59, 2},
	}

	templates.ExecuteTemplate(w, "Index", produtos)
}
