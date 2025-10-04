package controllers

import (
	"go-crud/models/products"
	"net/http"
	"strconv"
	"text/template"
)

const defaultRedirectCode = 301

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
	templates.ExecuteTemplate(w, "Index", products)

}

func Add(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Add", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		quantity, _ := strconv.Atoi(r.FormValue("quantity"))

		products.Insert(name, description, price, quantity)
		http.Redirect(w, r, "/", defaultRedirectCode)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	products.Delete(id)
	http.Redirect(w, r, "/", defaultRedirectCode)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := products.GetById(id)
	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		quantity, _ := strconv.Atoi(r.FormValue("quantity"))

		products.Update(id, name, description, price, quantity)
		http.Redirect(w, r, "/", defaultRedirectCode)
	}
}
