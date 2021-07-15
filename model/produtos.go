package model

import "go/config/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {

	dbConn := db.ConectaComPostgreSQL()
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
	return produtos
}
