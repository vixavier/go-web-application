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
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul", Preco: 39, Quantidade: 5},
		{"Tênis", "Confortável", 89, 10},
		{"Fone", "Bluetooth", 59, 2},
		{"Blusa", "Vermelha", 79, 5},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
