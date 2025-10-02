package database

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

/*
Configura a conexão com o Postgres
*/
func Connect() *sql.DB {
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
