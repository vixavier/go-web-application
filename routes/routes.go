package routes

import (
	"go/handler"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/new", handler.New)
	http.HandleFunc("/insert", handler.Insert)
	http.HandleFunc("/delete", handler.Delete)
	http.HandleFunc("/edit", handler.Edit)
}
