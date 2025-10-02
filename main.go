package main

import (
	"database/sql"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

/*
Configura a conexão com o Postgres
*/
func dbConnect() *sql.DB {
	/*
		Carrega as variáveis de ambiente utilizando a lib godotenv
	*/
	err := godotenv.Load()
	if err != nil {
		os.Exit(-1)
	}

	dbPass := os.Getenv("POSTGRES_PASS")
	connection := "user=postgres dbname=go_crud password=" + dbPass + " host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

/*
Carrega todos os templates que estão no diretório "templates" e que tenham a extensão ".html"
*/
var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	/*
		Atende a request da rota "/" através da função "index"
	*/
	http.HandleFunc("/", index)

	/*
		Inicia o servidor
	*/
	http.ListenAndServe(":8000", nil)
}

/*
Função que atende a request da rota "/".
O parâmetro "w" representa o ResponseWriter é responsável por escrever a resposta
O parâmetro "r" representa o Request, contendo os dados da requisição
*/
func index(w http.ResponseWriter, r *http.Request) {
	/*
		Abre a conexão com o Postgres e fecha no fim da função após executá-la com o comando "defer"
	*/
	db := dbConnect()

	/*
		Executa a query "SELECT * FROM products"
	*/
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	products := []Product{}

	for rows.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = rows.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

		products = append(products, product)
	}

	/*
		Executa o template "index.html" passando os dados de produtos como parâmetro ou "nil" se o template não executa nenhuma ação
	*/
	templates.ExecuteTemplate(w, "index.html", products)

	defer db.Close()
}
