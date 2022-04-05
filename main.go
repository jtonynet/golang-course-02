package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func connectaComBancoDeDados() *sql.DB {
	//docker run --name bd-postgres -e "POSTGRES_PASSWORD=postgres" -p 5432:5432 -v docker_volume:/var/lib/postgresql/data -d postgres
	conexao := "user=postgres dbname=postgres password=postgres host=localhost sslmode=disable"
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
	http.ListenAndServe(":8001", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// produtos := []Produto{
	// 	{1, "Camiseta", "Bem bonita", 29., 10},
	// 	{2, "Tenis", "Confortavel", 89., 10},
	// 	{3, "Fone", "Muito bom", 29., 10},
	// }

	db := connectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&nome, &descricao, &preco, &quantidade, &id)
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

	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
