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

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	dbConn := db.ConectaComPostgreSQL()
	defer dbConn.Close()

	insertProduct, err := dbConn.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertProduct.Exec(nome, descricao, preco, quantidade)
}

func DeletaProduto(id string) {
	db := db.ConectaComPostgreSQL()
	defer db.Close()

	deletarOProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarOProduto.Exec(id)
}
