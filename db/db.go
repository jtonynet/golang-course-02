package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectaComBancoDeDados() *sql.DB {
	/*
		// TODO: MOVE TO VOLUMES

		CREATE TABLE public.produtos (
			id serial,
			nome varchar NOT NULL,
			descricao varchar NOT NULL,
			preco decimal NOT NULL,
			quantidade int NOT NULL
		);

		INSERT INTO public.produtos (nome,descricao,preco,quantidade) VALUES
			('Camiseta','Preta',19,10),
			('Fone','Muito Bom',99,2),
			('Bermuda','teste mesmo',199,3);
	*/

	//docker run --name bd-postgres -e "POSTGRES_PASSWORD=postgres" -p 5432:5432 -v docker_volume:/var/lib/postgresql/data -d postgres

	conexao := "user=postgres dbname=postgres password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
