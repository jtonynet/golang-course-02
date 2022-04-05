package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectaComBancoDeDados() *sql.DB {
	//docker run --name bd-postgres -e "POSTGRES_PASSWORD=postgres" -p 5432:5432 -v docker_volume:/var/lib/postgresql/data -d postgres
	conexao := "user=postgres dbname=postgres password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
