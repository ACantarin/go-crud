package routes

import (
	"go-crud/controllers"
	"net/http"
)

func Load() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/add", controllers.Add)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
}
