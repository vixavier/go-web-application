package routes

import (
	"go/handler"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/new", handler.New)
}
