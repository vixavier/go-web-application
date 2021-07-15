package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComPostgreSQL() *sql.DB {
	conexao := "user=postgres dbname=golang_web_app_store password=pwdatabase host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
