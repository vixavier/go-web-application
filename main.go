package main

import (
	"go/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}
