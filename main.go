package main

import (
	"go-crud/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	/*
		Carrega as rotas
	*/
	routes.Load()

	/*
		Inicia o servidor
	*/
	http.ListenAndServe(":8000", nil)
}
