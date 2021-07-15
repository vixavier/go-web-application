package handler

import (
	"go/model"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := model.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", todosProdutos)
}
