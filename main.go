package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaComPostgreSQL() *sql.DB {
	conexao := "user=postgres dbname=golang_web_app_store password=pwdatabase host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Id         int
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
	// produtos := []Produto{
	// 	{Nome: "Camiseta", Descricao: "Azul", Preco: 39, Quantidade: 5},
	// 	{"Tênis", "Confortável", 89, 10},
	// }
	dbConn := conectaComPostgreSQL()
	getAllProducts, err := dbConn.Query("select * from produtos")
	defer dbConn.Close() // defer é executado apenas no final da função, após todas as outras linhas de comando

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for getAllProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = getAllProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
