package controllers

import (
	"net/http"
	"text/template"

	models "github.com/jtonynet/golang-course-02/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}
