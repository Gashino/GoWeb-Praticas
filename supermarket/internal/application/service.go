package application

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func StartRouter() {
	var router = chi.NewRouter()

	//Here starts the application
	AttachProductsController(router)

	//Start listening
	log.Fatal(http.ListenAndServe(":8080", router))
}
