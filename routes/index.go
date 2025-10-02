package routes

import (
	"go-crud/controllers"
	"net/http"
)

func Load() {
	http.HandleFunc("/", controllers.Index)
}
