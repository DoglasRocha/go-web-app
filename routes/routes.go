package routes

import (
	cont "loja/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", cont.Index)
	http.HandleFunc("/new", cont.New)
}
