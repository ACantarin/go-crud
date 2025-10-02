package controllers

import (
	"go-crud/models/products"
	"net/http"
	"text/template"
)

/*
Carrega todos os templates que estão no diretório "templates" e que tenham a extensão ".html"
*/
var templates = template.Must(template.ParseGlob("templates/*.html"))

/*
Função que atende a request da rota "/".
O parâmetro "w" representa o ResponseWriter é responsável por escrever a resposta
O parâmetro "r" representa o Request, contendo os dados da requisição
*/
func Index(w http.ResponseWriter, r *http.Request) {
	products := products.GetAll()

	/*
		Executa o template "index.html" passando os dados de produtos como parâmetro ou "nil" se o template não executa nenhuma ação
	*/
	templates.ExecuteTemplate(w, "index.html", products)

}
