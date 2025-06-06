package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	r := chi.NewRouter()

	r.Get("/ping", func(writer http.ResponseWriter, request *http.Request) {

		writer.WriteHeader(http.StatusOK)
		fmt.Fprint(writer, "pong")
	})

	r.Post("/greetings", func(writer http.ResponseWriter, request *http.Request) {
		var person Person
		err := json.NewDecoder(request.Body).Decode(&person)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Hola %s %s!", person.FirstName, person.LastName)
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}

}
