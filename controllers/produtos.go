package controllers

import (
	"html/template"
	"log"
	produtos "loja/models"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := produtos.BuscaTodosOsProdutos()

	templates.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		produtos.CriarNovoProduto(nome, descricao, preco, quantidade)
	}

	http.Redirect(w, r, "/", 301)
}
