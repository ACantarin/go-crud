package main

import (
	"net/http"
	"text/template"
)

type Product struct {
	Name  string
	Price float64
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
		Executa o template "index.html" passando o parâmetro "nil" pois o template não executa nenhuma ação
	*/
	templates.ExecuteTemplate(w, "index.html", nil)
}
