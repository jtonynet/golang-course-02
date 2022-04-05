package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8001", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Camiseta", "Bem bonita", 29., 10},
		{"Tenis", "Confortavel", 89., 10},
		{"Fone", "Muito bom", 29., 10},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
